package main

import (
	"log"
	"os"
	"time"

	"github.com/jakecallery/iiria/worker/cacheClient"
	"github.com/jakecallery/iiria/worker/keymaps"
	"github.com/jakecallery/iiria/worker/weatherClients"
)

type WeatherWorker struct {
	l        *log.Logger
	wcl      *weatherClients.ClientConfig
	ccl      cacheClient.CacheClient
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
	ww.get(ww.wcl) //Run first pull before timer to prime cache
	ww.t = time.NewTicker(5 * time.Minute)
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
	ww.l.Println("Tick")

	if os.Getenv(keymaps.EnvKeyMap[keymaps.LocalOnly]) == "true" {
		c.ExampleResponse = weatherClients.ExampleResponse
	}

	crd, err := c.Call()

	if err != nil {
		ww.l.Fatalf("[ERROR]: Error Calling API: %v", err)
	}

	ww.ccl.Save(crd)
}
