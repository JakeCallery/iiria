package cacheClient

import (
	"github.com/jakecallery/iiria/worker/weatherClients"
)

type CacheClient interface {
	Init()
	CheckConnection() error
	Save(*weatherClients.WeatherData) error
}
