package keymaps

import (
	"fmt"
)

type RangeDesc []string

func NewRangeDesc() RangeDesc {
	a := RangeDesc{
		"Low",
		"Low",
		"Low",
		"Moderate",
		"Moderate",
		"Moderate",
		"High",
		"High",
		"Very high",
		"Very high",
		"Very high",
		"Extreme",
	}

	return a

}

func (rd RangeDesc) GetDesc(idx int) (string, error) {
	if idx < 0 {
		return "", fmt.Errorf("ERROR: Invalid index (%v) for RangeDesc.  Valid values are %v - %v", idx, 0, len(rd)-1)
	}
	if idx > len(rd) {
		return rd[len(rd)-1], nil
	}

	return rd[idx], nil
}
