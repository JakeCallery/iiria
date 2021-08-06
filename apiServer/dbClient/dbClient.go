package dbClient

import "github.com/jakecallery/iiria/apiServer/data"

type DbClient interface {
	Init()
	DataFromTime(string) *data.WeatherData
	CheckConnection() error
}
