package http

import (
	"fmt"
	"log"
	"net/http"

	paho "github.com/eclipse/paho.mqtt.golang"
)

var (
	c paho.Client
)

func init() {
	opts := paho.NewClientOptions()
	opts.AddBroker("tcp://iot.eclipse.org:1883")
	opts.SetClientID("http")
	c = paho.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "It Works!") // nolint
}

func ligarLed(w http.ResponseWriter, r *http.Request) {
	// Publish a message.
	token := c.Publish("leds", 0, false, []byte("liga"))
	if token.Wait() && token.Error() != nil {
		http.Error(w, token.Error().Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func desligarLed(w http.ResponseWriter, r *http.Request) {
	// Publish a message.

	token := c.Publish("leds", 0, false, []byte("desliga"))
	if token.Wait() && token.Error() != nil {
		http.Error(w, token.Error().Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
