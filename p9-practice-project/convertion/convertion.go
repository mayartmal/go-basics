package convertion

import (
	"errors"
	"strconv"
)

func Strings2Floats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for il, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New(("ERROR DURING PARSING"))
		}
		floats[il] = floatPrice
	}
	return floats, nil
}
