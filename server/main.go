package main

/****** .env Example ******
apikey=<apikey>
latlong=lat,long
baseurl=https://api.tomorrow.io/v4/timelines?
localonly=true
****************************/

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/jakecallery/iiria/server/handlers"
	"github.com/jakecallery/iiria/server/keymaps"
	"github.com/jakecallery/iiria/server/weatherClients"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c := weatherClients.NewDefaultClientConfig()

	if os.Getenv(keymaps.EnvKeyMap[keymaps.LocalOnly]) == "true" {
		c.ExampleResponse = weatherClients.ExampleResponse
	}

	crd, err := c.Call()

	if err != nil {
		log.Fatalf("[ERROR]: Error Calling API: %v", err)
	}

	log.Printf("Time: %v", crd.Data.Timelines[0].Intervals[0].StartTime)
	log.Printf("Temp: %v", crd.Data.Timelines[0].Intervals[0].Values.Temperature)
	log.Printf("PrecipType: %v", keymaps.PrecipTypeCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.PrecipitationType)])
	log.Printf("WeatherCode: %v", keymaps.WeatherCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.WeatherCode)])

	rd := keymaps.NewRangeDesc()
	uvIndex, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVIndex)

	if err != nil {
		log.Printf("[ERROR]: Failed to retrieve a valid uvIndex: %v\n", err)
		log.Printf("[ERROR]: Setting to 'unknown'")
		uvIndex = "Unknown"
	}

	uvHealth, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVHealthConcern)
	if err != nil {
		log.Printf("[ERROR]: Failed to retrieve a valid uvHealth Concern: %v", err)
		log.Printf("[ERROR]: setting to 'Unknown'")
		uvHealth = "Unknown"
	}

	log.Printf("UVIndex: %v", uvIndex)
	log.Printf("UVHealthConcern: %v", uvHealth)

	l := log.New(os.Stdout, "[weather-api]", log.LstdFlags)
	wh := handlers.NewCurrentWeather(l)
	hh := handlers.NewHealth(l)
	sm := http.NewServeMux()
	sm.Handle("/", wh)
	sm.Handle("/health", hh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
