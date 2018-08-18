package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// StartServer http
func StartServer() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: makeHandlers(),
	}
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		log.Println("Stopping server http ...")
		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}()
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
		return
	}
}

func makeHandlers() http.Handler {
	n := negroni.Classic()
	n.Use(basicAuth())
	n.UseHandler(makeRoutes())
	return n
}

func basicAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if r.URL.Path != "/" {
			user, pass, ok := r.BasicAuth()
			if !ok || (user != "iot" && pass != "iot") {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintln(w, http.StatusText(http.StatusUnauthorized)) // nolint
				return
			}
		}
		next(w, r)
	})
}

func makeRoutes() http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", home)
	r.HandleFunc("/led/ligar", ligarLed).Methods(http.MethodPut)
	r.HandleFunc("/led/desligar", desligarLed).Methods(http.MethodPut)
	return r
}
