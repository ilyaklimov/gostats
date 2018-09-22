package gostats

import (
	"errors"
)

var errDataIsEmpty          = errors.New("data is empty")
var errInsufficientData     = errors.New("insufficient data (len(data) == 1)")
