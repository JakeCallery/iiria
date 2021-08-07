package dataGetter

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

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

func (dg *DataGetter) GetData() (*data.WeatherData, error) {
	dg.l.Println("Get me some data!")
	dg.l.Printf("Time: %s", getServerTime())
	res, err := dg.db.DataFromTime(getServerTime())

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

func getServerTime() string {
	ti := time.Now().UTC()

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(ti.Year()))
	sb.WriteString("-")
	sb.WriteString(fmt.Sprintf("%02d", ti.Month()))
	sb.WriteString("-")
	sb.WriteString(fmt.Sprintf("%02d", ti.Day()))
	sb.WriteString("T")
	sb.WriteString(fmt.Sprintf("%02d", ti.Hour()))
	sb.WriteString("_")
	sb.WriteString(fmt.Sprintf("%02d", ti.Minute()))
	sb.WriteString("_00Z")
	return sb.String()

}
