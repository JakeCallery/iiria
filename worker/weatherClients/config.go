package weatherClients

import (
	"os"

	"github.com/jakecallery/iiria/worker/keymaps"
)

type ClientConfig struct {
	BaseURL         string
	ApiKey          string
	LatLong         string
	Fields          []string
	TimeSteps       []string
	Timezone        string
	ExampleResponse []byte
}
func NewDefaultClientConfig() *ClientConfig {

	return NewClientConfig(
		os.Getenv(keymaps.EnvKeyMap[keymaps.BaseURL]),
		os.Getenv(keymaps.EnvKeyMap[keymaps.APIkey]),
		os.Getenv(keymaps.EnvKeyMap[keymaps.LatLong]),
		[]string{
			"temperature",
			"precipitationType",
			"weatherCode",
			"uvIndex",
			"uvHealthConcern",
		},
		[]string{"1m"},
		"America/New_York",
	)
}

func NewClientConfig(
	baseURL string,
	apiKey string,
	latLong string,
	fields []string,
	timeSteps []string,
	timeZone string,

) *ClientConfig {

	cc := new(ClientConfig)

	cc.BaseURL = baseURL
	cc.ApiKey = apiKey
	cc.LatLong = latLong
	cc.Fields = fields
	cc.TimeSteps = timeSteps
	cc.Timezone = timeZone

	return cc

}
