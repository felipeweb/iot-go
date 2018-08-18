package iot

import (
	"log"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
	"gobot.io/x/gobot/platforms/mqtt"
)

var (
	led1        *gpio.LedDriver
	led2        *gpio.LedDriver
	mqttAdaptor *mqtt.Adaptor
)

func work() {
	mqttAdaptor.On("leds", func(msg mqtt.Message) {
		if string(msg.Payload()) == "liga" {
			err := led1.On()
			if err != nil {
				log.Fatal(err)
			}
			err = led2.On()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err := led1.Off()
			if err != nil {
				log.Fatal(err)
			}
			err = led2.Off()
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

// Setup arduino
func Setup() {
	mqttAdaptor = mqtt.NewAdaptor("tcp://iot.eclipse.org:1883", "arduino")
	firmataAdaptor := firmata.NewAdaptor("/dev/cu.usbmodem1411")
	led1 = gpio.NewLedDriver(firmataAdaptor, "10")
	led2 = gpio.NewLedDriver(firmataAdaptor, "13")
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor, mqttAdaptor},
		[]gobot.Device{led1, led2},
		work,
	)
	go func() {
		err := robot.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
