package main

/****** .env Example ******
apikey=<apikey>
latlong=lat,long
baseurl=https://api.tomorrow.io/v4/timelines?
localonly=true
****************************/

import (
	"log"
	"os"
	"strconv"

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

}
