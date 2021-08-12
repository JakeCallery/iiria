package data

import (
	"encoding/json"
	"log"
)

type WeatherData struct {
	Temperature       float64 `json:"temperature"`
	PrecipitationType int     `json:"precipitationType"`
	PrecipitationDesc string  `json:"precipitationDesc"`
	WeatherCode       int     `json:"weatherCode"`
	WeatherDesc       string  `json:"weatherDesc"`
	UvIndex           int     `json:"uvIndex"`
	UvDesc            string  `json:"uvDesc"`
	UvHealthConcern   int     `json:"uvHealthConcern"`
	UvHealthDesc      string  `json:"uvHealthDesc"`
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

	desc, err := GetWeatherDescFromCode(wd.WeatherCode)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return err
	}
	wd.WeatherDesc = desc

	desc, err = GetPrecipTypeFromCode(wd.PrecipitationType)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return err
	}
	wd.PrecipitationDesc = desc

	rd := NewRangeDesc()

	desc, err = rd.GetDesc(wd.UvIndex)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return err
	}
	wd.UvDesc = desc

	desc, err = rd.GetDesc(wd.UvHealthConcern)
	if err != nil {
		l.Printf("[ERROR]: %v", err)
		return err
	}
	wd.UvHealthDesc = desc

	return nil
}
