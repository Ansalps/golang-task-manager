package utils

import (
	"errors"
	"strconv"
)

func StringConversion(s string) (uint, error) {
	intValue, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("error converting string")
	}
	return uint(intValue), nil
}
