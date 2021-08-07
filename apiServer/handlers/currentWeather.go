package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jakecallery/iiria/apiServer/dataGetter"
)

type CurrentWeather struct {
	l  *log.Logger
	dg *dataGetter.DataGetter
}

func NewCurrentWeather(l *log.Logger, dg *dataGetter.DataGetter) *CurrentWeather {
	return &CurrentWeather{l, dg}
}

func (h *CurrentWeather) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Data Requested")
	wd, err := h.dg.GetData()

	//TODO: Report proper error based on returned error.
	//For now just returning a 500

	if err != nil {
		h.l.Printf("[ERROR] Error Getting Data: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	respData, err := json.Marshal(wd)
	if err != nil {
		h.l.Printf("[ERROR] Error converting response data to json: %v", err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, "%s", respData)
}
