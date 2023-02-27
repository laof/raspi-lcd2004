package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	device "github.com/d2r2/go-hd44780"
	i2c "github.com/d2r2/go-i2c"
	logger "github.com/d2r2/go-logger"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func d2(nub float64) float64 {
	nub, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", nub), 64)
	return nub
}

var mutex sync.Mutex

func main() {

	logger.ChangePackageLogLevel("i2c", logger.InfoLevel)
	i2c, err := i2c.NewI2C(0x27, 1)
	check(err)

	lcd, e := device.NewLcd(i2c, device.LCD_20x4)
	check(e)
	lcd.BacklightOn()
	lcd.Clear()

	defer lcd.Clear()
	defer i2c.Close()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " LCD_20x4 ok")

	lcd.Home()
	// welcome
	lcd.SetPosition(0, 0)
	fmt.Fprint(lcd, "")

	lcd.SetPosition(1, 0)
	fmt.Fprint(lcd, "")

	go func() {
		for {
			info(lcd)
			time.Sleep(3 * time.Second)
		}
	}()

	for {
		mutex.Lock()
		t := time.Now()
		lcd.SetPosition(2, 0)
		fmt.Fprint(lcd, t.Format("Monday Jan"))
		lcd.SetPosition(3, 0)
		fmt.Fprint(lcd, t.Format("15:04:05  2006-01-02"))
		mutex.Unlock()
		time.Sleep(1 * time.Second)
	}
}

func info(lcd *device.Lcd) {

	cmd := exec.Command("vcgencmd", "measure_temp")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	tempByte, _ := ioutil.ReadAll(&out)

	temp := string(tempByte)

	temp = strings.Replace(temp, "temp=", "", 1)

	temp = strings.Split(string(temp), "'C")[0]

	idle0, total0 := getCPUSample()
	time.Sleep(3 * time.Second)
	idle1, total1 := getCPUSample()

	idleTicks := d2(float64(idle1 - idle0))
	totalTicks := d2(float64(total1 - total0))
	cpuUsage := d2(100 * (totalTicks - idleTicks) / totalTicks)

	network := NetWorkStatus()
	base := fmt.Sprintf("CPU:%v%v %v'C   ", cpuUsage, "%", temp)

	mutex.Lock()
	lcd.SetPosition(0, 0)
	fmt.Fprint(lcd, "Wifi network: "+network)
	lcd.SetPosition(1, 0)
	fmt.Fprint(lcd, base)
	mutex.Unlock()

}

func getCPUSample() (idle, total uint64) {
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

func NetWorkStatus() string {
	cmd := exec.Command("ping", "baidu.com", "-c", "1", "-W", "5")
	err := cmd.Run()
	if err == nil {
		return "ok"
	}
	return "no"
}
