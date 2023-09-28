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
	"syscall"
	"time"

	"github.com/jakecallery/iiria/apiServer/dataGetter"
	"github.com/jakecallery/iiria/apiServer/dbClient"
	"github.com/jakecallery/iiria/apiServer/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	l := log.New(os.Stdout, "[weather-api]", log.LstdFlags)

	//Set up redis connection
	rh := os.Getenv("redishost")
	rp := os.Getenv("redisport")

	if rh == "" {
		rh = "localhost"
	}

	if rp == "" {
		rp = "6379"
	}

	db := dbClient.NewRedisClient(l, rh, rp)
	db.Init()
	db.CheckConnection()

	dg := dataGetter.NewDataGetter(l, db)

	wh := handlers.NewCurrentWeather(l, dg)
	hh := handlers.NewHealth(l)

	sm := http.NewServeMux()
	sm.Handle("/api/weather", wh)
	sm.Handle("/api/health", hh)

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

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	tc, tcCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer tcCancel()
	s.Shutdown(tc)
}
