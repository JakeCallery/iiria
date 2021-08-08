package data

import (
	"errors"
	"strconv"
)

var PrecipTypeCodes map[string]string = map[string]string{
	"0": "N/A",
	"1": "Rain",
	"2": "Snow",
	"3": "Freezing Rain",
	"4": "Ice Pellets",
}

func GetPrecipTypeFromCode(c int) (string, error) {
	str := strconv.Itoa(c)
	if desc, ok := PrecipTypeCodes[str]; ok {
		return desc, nil
	}

	return "", errors.New("Bad Precip Type code: " + strconv.Itoa(c))
}
