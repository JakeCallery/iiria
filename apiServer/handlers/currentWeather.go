package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CurrentWeather struct {
	l *log.Logger
}

func NewCurrentWeather(l *log.Logger) *CurrentWeather {
	return &CurrentWeather{l}
}

func (h *CurrentWeather) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Printf("Request Error: %v", err)
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Data: %s\n", d)
}
