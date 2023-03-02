package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	device "github.com/d2r2/go-hd44780"
	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
)

//go:embed api
var restapi string

//go:embed en
var english string

var list []string = strings.Split(english, "\r\n")

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func d2(nub float64) float64 {
	nub, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", nub), 64)
	return nub
}

func show(line int, txt string) {
	mutex.Lock()
	lcd.SetPosition(line, 0)
	fmt.Fprint(lcd, safeScreen(txt))
	mutex.Unlock()
}

var (
	mutex  sync.Mutex
	failed = "Failed"
)

var lcd *device.Lcd

func main() {

	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	i2c, err := i2c.NewI2C(0x27, 1)
	check(err)

	lcd, _ = device.NewLcd(i2c, device.LCD_20x4)

	lcd.BacklightOff()
	lcd.Clear()

	defer lcd.Clear()
	defer i2c.Close()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " LCD_20x4 ok")

	lcd.Home()
	// wifi
	show(0, "Hello,")

	// cpu
	show(1, "Raspberry Pi!")

	//weather
	show(2, "")

	go func() {
		for {
			wifiCpuInfo()
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		time.Sleep(10 * time.Second)
		weatherInfo()
		for {
			time.Sleep(30 * time.Minute)
			weatherInfo()
		}
	}()

	for {
		t := time.Now()
		show(3, t.Format("15:04:05  2006-01-02"))
		time.Sleep(1 * time.Second)
	}
}

func wifiCpuInfo() {

	idle0, total0 := cpu()
	time.Sleep(3 * time.Second)
	idle1, total1 := cpu()

	idleTicks := d2(float64(idle1 - idle0))
	totalTicks := d2(float64(total1 - total0))
	cpuUsage := d2(100 * (totalTicks - idleTicks) / totalTicks)

	base := fmt.Sprintf("CPU: %v%v %s", cpuUsage, "%", temp())

	net := network()

	show(0, "Network: "+net)
	show(1, base)

}

func cpu() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func network() string {
	var out bytes.Buffer
	cmd := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5")
	cmd.Stdout = &out
	err := cmd.Run()

	if err == nil {
		tb, _ := io.ReadAll(&out)
		output := string(tb)
		reg, _ := regexp.Compile(`time=(.*)\n`)
		time := reg.FindStringSubmatch(output)

		if len(time) > 1 {
			return time[1]
		}
	}
	return failed
}

func temp() (output string) {
	var out bytes.Buffer

	cmd := exec.Command("vcgencmd", "measure_temp")
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	tb, _ := io.ReadAll(&out)
	output = string(tb)

	reg, _ := regexp.Compile(`temp=(.*)`)
	tp := reg.FindStringSubmatch(output)

	if len(tp) > 1 {
		output = tp[1]
	} else {
		output = failed
	}
	return
}

var empty = strings.Repeat(" ", 20)

func safeScreen(txt string) string {
	max := 20
	num := len(txt) - max

	if num > 0 {
		return fmt.Sprintf("%v%v", txt[0:18], "..")
	}

	num = int(math.Abs(float64(num)))

	return fmt.Sprintf("%v%v", txt, empty[0:num])

}

func weatherInfo() {

	w, t := weatherapi()

	var en string = "unknown"

	for i, v := range list {

		if v == w {
			en = list[i+1]
			break
		}
	}

	hour := time.Now().Format("15")
	h, _ := strconv.Atoi(hour)

	if strings.Contains(w, "雨") && h >= 7 && h < 21 {
		lcd.BacklightOn()
	} else {
		lcd.BacklightOff()
	}

	show(2, fmt.Sprintf("%v %v'C", en, t))

}

type JSONData struct {
	Status   string  `json:"status"`
	Count    string  `json:"count"`
	Info     string  `json:"info"`
	Infocode string  `json:"infocode"`
	Lives    []Lives `json:"lives"`
}
type Lives struct {
	Province         string `json:"province"`
	City             string `json:"city"`
	Adcode           string `json:"adcode"`
	Weather          string `json:"weather"`
	Temperature      string `json:"temperature"`
	Winddirection    string `json:"winddirection"`
	Windpower        string `json:"windpower"`
	Humidity         string `json:"humidity"`
	Reporttime       string `json:"reporttime"`
	TemperatureFloat string `json:"temperature_float"`
	HumidityFloat    string `json:"humidity_float"`
}

func weatherapi() (string, string) {

	resr, err := http.Get(restapi)

	if err != nil {
		return "-", "-"
	}

	defer resr.Body.Close()

	body, e := io.ReadAll(resr.Body)

	if e != nil {
		return "-", "-"
	}

	var data *JSONData
	err = json.Unmarshal(body, &data)

	if err != nil {
		return "-", "-"
	}

	return data.Lives[0].Weather, data.Lives[0].Temperature

}
