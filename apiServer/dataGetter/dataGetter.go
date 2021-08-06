package datagetter

import (
	"log"

	"github.com/jakecallery/iiria/apiServer/data"
	"github.com/jakecallery/iiria/apiServer/dbClient"
)

type DataGetter struct {
	l  *log.Logger
	db *dbClient.DbClient
}

func NewDataGetter(l *log.Logger, db *dbClient.DbClient) *DataGetter {
	dg := &DataGetter{l, db}

	return dg
}

func (dg *DataGetter) getData() *data.WeatherData {
	dg.l.Panicln("Get me some data!")
	return nil
}
