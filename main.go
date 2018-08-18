package main

import (
	"github.com/felipeweb/iot-go/http"
	"github.com/felipeweb/iot-go/iot"
)

func main() {
	iot.Setup()
	http.StartServer()
}
