package weatherAPIClient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const baseURL string = "https://api.tomorrow.io/v4/timelines?"

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

type CurrentWeatherGetClient struct {
	ApiKey    string
	LatLong   string
	Fields    []string
	TimeSteps []string
	Timezone  string
}

//TODO: Research this: https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

func (c *CurrentWeatherGetClient) Call() /*(*http.Response, error)*/ {
	var sb strings.Builder
	sb.WriteString(baseURL)
	sb.WriteString("location=" + c.LatLong)
	sb.WriteString("&fields=" + strings.Join(c.Fields[:], ","))
	sb.WriteString("&timesteps=" + strings.Join(c.TimeSteps[:], ","))
	sb.WriteString("&units=imperial")
	sb.WriteString("&timezone=" + c.Timezone)
	sb.WriteString("&apikey=" + c.ApiKey)

	fmt.Printf("\nURL: " + sb.String() + "\n")

	resp, err := netClient.Get(sb.String())

	if err != nil {
		log.Fatalf("Request Error: %+v", err)
	}

	defer resp.Body.Close()

	var jsonData map[string]interface{}
	bodyData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Body Read Error: %+v", err)
	}

	err = json.Unmarshal([]byte(bodyData), &jsonData)

	if err != nil {
		log.Fatalf("Json Parse Error: %+v", err)
	}

	//fmt.Println(jsonData)
	fmt.Println(string(bodyData))
}
