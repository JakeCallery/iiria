package handlers

import (
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
	// d, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	h.l.Printf("Request Error: %v", err)
	// 	http.Error(rw, "oops", http.StatusBadRequest)
	// 	return
	// }
	wd, _ := h.dg.GetData()
	fmt.Fprintf(rw, "Data: %+v\n", wd)
}
