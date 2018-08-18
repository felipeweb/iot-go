package iot

import (
	"log"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

var (
	led1 *gpio.LedDriver
	led2 *gpio.LedDriver
)

func init() {
	setup()
}

// GetLed1 obj
func GetLed1() *gpio.LedDriver {
	return led1
}

// GetLed2 obj
func GetLed2() *gpio.LedDriver {
	return led2
}

// Setup arduino
func setup() {
	firmataAdaptor := firmata.NewAdaptor("/dev/cu.usbmodem1411")
	led1 = gpio.NewLedDriver(firmataAdaptor, "10")
	led2 = gpio.NewLedDriver(firmataAdaptor, "13")
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led1, led2},
	)
	go func() {
		err := robot.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
