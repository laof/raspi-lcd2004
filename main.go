package main

import (
	"fmt"
	"log"
	"time"

	i2c "github.com/laof/go-i2c"
	device "github.com/laof/go-lcd"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i2c, err := i2c.NewI2C(0x27, 1)
	check(err)

	lcd, err := device.NewLcd(i2c, device.LCD_20x4)
	check(err)
	lcd.BacklightOff()
	lcd.Clear()

	defer lcd.Clear()
	defer i2c.Close()

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), " LCD_20x4 ok")
	// welcome
	lcd.SetPosition(0, 0)
	fmt.Fprint(lcd, "Hi, Raspberry Pi 4b!")
	lcd.SetPosition(1, 0)
	fmt.Fprint(lcd, time.Now().Format("01-02 15:04:05"))
	for {
		lcd.Home()
		t := time.Now()
		lcd.SetPosition(2, 0)
		fmt.Fprint(lcd, t.Format("Monday Jan 2"))
		lcd.SetPosition(3, 0)
		fmt.Fprint(lcd, t.Format("15:04:05  2006-01-02"))
		time.Sleep(3 * time.Second)
	}
}
