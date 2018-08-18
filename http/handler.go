package http

import (
	"fmt"
	"net/http"

	"github.com/felipeweb/iot-go/iot"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It Works!") // nolint
}

func ligarLed1(w http.ResponseWriter, r *http.Request) {
	led := iot.GetLed1()
	err := led.On()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ligarLed2(w http.ResponseWriter, r *http.Request) {
	led := iot.GetLed2()
	err := led.On()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func desligarLed1(w http.ResponseWriter, r *http.Request) {
	led := iot.GetLed1()
	err := led.Off()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func desligarLed2(w http.ResponseWriter, r *http.Request) {
	led := iot.GetLed2()
	err := led.Off()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
