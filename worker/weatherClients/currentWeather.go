package weatherClients

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jakecallery/iiria/server/keymaps"
)

func (c *clientConfig) Call() (*CurrentResponseData, error) {

	var body []byte
	var err error

	if c.ExampleResponse == nil {
		log.Println("Going to internet to get data...")
		body, err = getData(c)

		if err != nil {
			log.Printf("[ERROR]: %v\n", err)
			return nil, err
		}

	} else {
		log.Println("--- Using Example Data ---")
		body = []byte(c.ExampleResponse)
	}

	crd := CurrentResponseData{}
	err = jsonToStruct(body, &crd)

	if err != nil {
		log.Printf("[ERROR]: %v\n", err)
		return nil, err
	}

	//log.Printf("Struct: \n%+v", crd)

	return &crd, nil

}

func buildURL(c *clientConfig) string {
	var sb strings.Builder
	sb.WriteString(os.Getenv(keymaps.EnvKeyMap[keymaps.BaseURL]))
	sb.WriteString("location=" + c.LatLong)
	sb.WriteString("&fields=" + strings.Join(c.Fields[:], ","))
	sb.WriteString("&timesteps=" + strings.Join(c.TimeSteps[:], ","))
	sb.WriteString("&units=imperial")
	sb.WriteString("&timezone=" + c.Timezone)
	sb.WriteString("&apikey=" + c.ApiKey)

	log.Printf("\nURL: " + sb.String() + "\n")
	return sb.String()
}

func jsonToStruct(d []byte, crd *CurrentResponseData) error {
	err := json.Unmarshal(d, &crd)

	if err != nil {
		log.Fatalf("Failed to marshal json: %v", err)
		return err
	}

	return nil

}

func getData(c *clientConfig) ([]byte, error) {
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

	return body, nil
}
