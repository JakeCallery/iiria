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
	"syscall"
	"time"

	"github.com/jakecallery/iiria/worker/cacheClient"

	"github.com/jakecallery/iiria/worker/weatherClients"
	"github.com/joho/godotenv"
)

func main() {


	l := log.New(os.Stdout, "[WorkerMain]: ", log.LstdFlags)

	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cacheClient := cacheClient.NewRedisClient(log.New(os.Stdout, "[cacheClient]: ", log.LstdFlags))
	cacheClient.Init()
	err = cacheClient.CheckConnection()

	if err != nil {
		l.Fatalf("Failed to get a good cache server connection: %v", err)
	}

	stopChan := make(chan bool)
	c := weatherClients.NewDefaultClientConfig()
	ww := NewWeatherWorker(log.New(os.Stdout, "[WeatherWorker]: ", log.LstdFlags), c, cacheClient, stopChan)
	go func() {
		ww.Run()
	}()

	//Shutdown handling
	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	_, tcCancel := context.WithTimeout(context.Background(), 30*time.Second)
	ww.Stop()
	defer tcCancel()
}
