package main

/****** .env Example ******
apikey=<apikey>
latlong=lat,long
baseurl=https://api.tomorrow.io/v4/timelines?
****************************/

import (
	"log"

	//"weatherAPIClient"

	"github.com/jakecallery/iiria/server/weatherClients"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	c := weatherClients.NewDefaultClientConfig()
	c.Call()

}
