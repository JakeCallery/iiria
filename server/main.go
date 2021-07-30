package main

import (
	"fmt"
	"log"
	"os"

	//"weatherAPIClient"
	k "github.com/jakecallery/iiria/server/keymaps"
	"github.com/jakecallery/iiria/server/weatherAPIClient"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Printf("Key: %v\n", os.Getenv(k.EnvKeyMap[k.APIkey]))

	c := weatherAPIClient.CurrentWeatherGetClient{
		ApiKey:  os.Getenv(k.EnvKeyMap[k.APIkey]),
		LatLong: "30.6727578,-97.8365732",
		Fields: []string{
			"temperature",
			"precipitationType",
			"weatherCode",
		},
		TimeSteps: []string{"1m"},
		Timezone:  "America/New_York",
	}

	//resp, err := c.Call()

	c.CallExample()

}
