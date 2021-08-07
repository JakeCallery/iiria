package dataGetter

import (
	"log"

	"github.com/jakecallery/iiria/apiServer/data"
	"github.com/jakecallery/iiria/apiServer/dbClient"
)

type DataGetter struct {
	l  *log.Logger
	db dbClient.DbClient
}

func NewDataGetter(l *log.Logger, db dbClient.DbClient) *DataGetter {
	dg := &DataGetter{l, db}

	return dg
}

func (dg DataGetter) GetData() (*data.WeatherData, error) {
	dg.l.Println("Get me some data!")
	res, err := dg.db.DataFromTime("2021-08-01T15_51_00-04_00")

	if err != nil {
		dg.l.Printf("[ERROR]: Error in GetData: %v", err)
		return nil, err
	}

	d := data.NewWeatherData()
	err = d.FillFromJson(dg.l, res)

	if err != nil {
		dg.l.Printf("[ERROR]: Failed to fill weather data from json response: %v", err)
		return nil, err
	}
	dg.l.Printf("***Result: %v", d.Temperature)

	return d, nil
}
