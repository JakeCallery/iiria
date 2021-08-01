package weatherClients

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jakecallery/iiria/server/keymaps"
)

func (c *clientConfig) Call() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(buildURL(c))

	if err != nil {
		log.Fatalf("Current Weather Get Failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to ready body from response: %v", err)
	}

	fmt.Println(string(body))
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

	fmt.Printf("\nURL: " + sb.String() + "\n")
	return sb.String()
}
