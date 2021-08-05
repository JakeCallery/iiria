package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jakecallery/iiria/worker/cacheClient"
	"github.com/jakecallery/iiria/worker/keymaps"
	"github.com/jakecallery/iiria/worker/weatherClients"
)

//TODO: Set up cache client reference here
type WeatherWorker struct {
	l        *log.Logger
	wcl      *weatherClients.ClientConfig
	ccl      cacheClient.CacheClient //TODO: Use the interface CacheClient instead
	t        *time.Ticker
	stopChan chan bool
}

func NewWeatherWorker(l *log.Logger, wcl *weatherClients.ClientConfig, ccl cacheClient.CacheClient, stopChan chan bool) *WeatherWorker {
	return &WeatherWorker{
		l,
		wcl,
		ccl,
		nil,
		stopChan,
	}
}

func (ww *WeatherWorker) Run() {
	ww.t = time.NewTicker(time.Second)
	defer ww.t.Stop()
	for {
		select {
		case <-ww.t.C:
			ww.get(ww.wcl)
		case stop := <-ww.stopChan:
			if stop {
				ww.l.Println("Stopping")
				return
			}
		}
	}
}

func (ww *WeatherWorker) Stop() {
	ww.l.Println("Stopping Run")
	ww.stopChan <- true
}

func (ww *WeatherWorker) get(c *weatherClients.ClientConfig) {
	//TODO: store in cache
	ww.l.Println("Tick")

	if os.Getenv(keymaps.EnvKeyMap[keymaps.LocalOnly]) == "true" {
		c.ExampleResponse = weatherClients.ExampleResponse
	}

	crd, err := c.Call()

	if err != nil {
		ww.l.Fatalf("[ERROR]: Error Calling API: %v", err)
	}

	ww.l.Printf("Time: %v", crd.Data.Timelines[0].Intervals[0].StartTime)
	ww.l.Printf("Temp: %v", crd.Data.Timelines[0].Intervals[0].Values.Temperature)
	ww.l.Printf("PrecipType: %v", keymaps.PrecipTypeCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.PrecipitationType)])
	ww.l.Printf("WeatherCode: %v", keymaps.WeatherCodes[strconv.Itoa(crd.Data.Timelines[0].Intervals[0].Values.WeatherCode)])

	rd := keymaps.NewRangeDesc()
	uvIndex, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVIndex)

	if err != nil {
		ww.l.Printf("[ERROR]: Failed to retrieve a valid uvIndex: %v\n", err)
		ww.l.Printf("[ERROR]: Setting to 'unknown'")
		uvIndex = "Unknown"
	}

	uvHealth, err := rd.GetDesc(crd.Data.Timelines[0].Intervals[0].Values.UVHealthConcern)
	if err != nil {
		ww.l.Printf("[ERROR]: Failed to retrieve a valid uvHealth Concern: %v", err)
		ww.l.Printf("[ERROR]: setting to 'Unknown'")
		uvHealth = "Unknown"
	}

	ww.l.Printf("UVIndex: %v", uvIndex)
	ww.l.Printf("UVHealthConcern: %v", uvHealth)

	ww.ccl.Save(crd)
}
