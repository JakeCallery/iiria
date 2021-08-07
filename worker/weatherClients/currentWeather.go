package weatherClients

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jakecallery/iiria/worker/keymaps"
)

func (c *ClientConfig) Call() (*WeatherData, error) {

	var body []byte
	var err error

	if c.ExampleResponse == nil {

		body, err = getData(c)

		if err != nil {
			log.Printf("[ERROR]: %v\n", err)
			return nil, err
		}

	} else {
		log.Println("--- Using Example Data ---")
		body = []byte(c.ExampleResponse)
	}

	crd := WeatherData{}

	err = jsonToStruct(body, &crd)

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		return nil, err
	}

	return &crd, nil

}

func buildURL(c *ClientConfig) string {

	var sb strings.Builder
	sb.WriteString(os.Getenv(keymaps.EnvKeyMap[keymaps.BaseURL]))
	sb.WriteString("location=" + c.LatLong)
	sb.WriteString("&fields=" + strings.Join(c.Fields[:], ","))
	sb.WriteString("&timesteps=" + strings.Join(c.TimeSteps[:], ","))
	sb.WriteString("&units=imperial")
	sb.WriteString("&timezone=" + c.Timezone)
	sb.WriteString("&apikey=" + c.ApiKey)

	return sb.String()
}

func jsonToStruct(d []byte, crd *WeatherData) error {
	err := json.Unmarshal(d, &crd)

	if err != nil {
		log.Fatalf("Failed to unmarshal json: %v", err)

		return err
	}

	return nil

}

func getData(c *ClientConfig) ([]byte, error) {

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(buildURL(c))

	if err != nil {
		log.Printf("Current Weather Get Failed: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to ready body from response: %v\n", err)
		return nil, err
	}

	log.Printf("Ratelimit-Limit: %v", resp.Header["Ratelimit-Limit"])
	log.Printf("Ratelimit-Reset: %v", resp.Header["Ratelimit-Reset"])
	log.Printf("Ratelimit-Remaining: %v", resp.Header["Ratelimit-Remaining"])
	log.Printf("X-Ratelimit-Limit-Day: %v", resp.Header["X-Ratelimit-Limit-Day"])
	log.Printf("X-Ratelimit-Limit-Hour: %v", resp.Header["X-Ratelimit-Limit-Hour"])
	log.Printf("X-Ratelimit-Limit-Second: %v", resp.Header["X-Ratelimit-Limit-Second"])
	log.Printf("X-Ratelimit-Remaining-Day: %v", resp.Header["X-Ratelimit-Remaining-Day"])
	log.Printf("X-Ratelimit-Remaining-Hour: %v", resp.Header["X-Ratelimit-Remaining-Hour"])
	log.Printf("X-Ratelimit-Remaining-Second: %v", resp.Header["X-Ratelimit-Remaining-Second"])

	return body, nil
}
