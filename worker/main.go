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
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/jakecallery/iiria/worker/cacheClient"
	"github.com/jakecallery/iiria/worker/keymaps"
	"github.com/jakecallery/iiria/worker/weatherClients"
	"github.com/joho/godotenv"
)

func main() {

	l := log.New(os.Stdout, "[WorkerMain]: ", log.LstdFlags)

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c := weatherClients.NewDefaultClientConfig()

	if os.Getenv(keymaps.EnvKeyMap[keymaps.LocalOnly]) == "true" {
		c.ExampleResponse = weatherClients.ExampleResponse
	}

	crd, err := c.Call()

	if err != nil {
		l.Fatalf("[ERROR]: Error Calling API: %v", err)
	}

	l.Printf("Time: %v", crd.Data.Timelines[0].Intervals[0].StartTime)
	l.Printf("Temp: %v", crd.Data.Timelines[0].Intervals[0].Values.Temperature)
	l.Printf("PrecipType: %v", keymaps.PrecipTypeCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.PrecipitationType)])
	l.Printf("WeatherCode: %v", keymaps.WeatherCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.WeatherCode)])

	rd := keymaps.NewRangeDesc()
	uvIndex, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVIndex)

	if err != nil {
		l.Printf("[ERROR]: Failed to retrieve a valid uvIndex: %v\n", err)
		l.Printf("[ERROR]: Setting to 'unknown'")
		uvIndex = "Unknown"
	}

	uvHealth, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVHealthConcern)
	if err != nil {
		l.Printf("[ERROR]: Failed to retrieve a valid uvHealth Concern: %v", err)
		l.Printf("[ERROR]: setting to 'Unknown'")
		uvHealth = "Unknown"
	}

	l.Printf("UVIndex: %v", uvIndex)
	l.Printf("UVHealthConcern: %v", uvHealth)

	cacheClient := cacheClient.NewRedisClient(log.New(os.Stdout, "[cacheClient]: ", log.LstdFlags))
	cacheClient.Init()
	err = cacheClient.CheckConnection()

	if err != nil {
		l.Fatalf("Failed to get a good cache server connection: %v", err)
	}

	//Shutdown handling
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	_, tcCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer tcCancel()

}
