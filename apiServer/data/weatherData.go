package data

import (
	"encoding/json"
	"log"
)

type WeatherData struct {
	Temperature       float64 `json:"temperature"`
	PrecipitationType int     `json:"precipitationType"`
	WeatherCode       int     `json:"weatherCode"`
	UvIndex           int     `json:"uvIndex"`
	UvHealthConcern   int     `json:"uvHealthConcern"`
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (wd *WeatherData) FillFromJson(l *log.Logger, data string) error {
	err := json.Unmarshal([]byte(data), wd)
	if err != nil {
		l.Printf("[ERROR]: Error converting json to WeatherData: %v", err)
		return err
	}

	return nil
}
