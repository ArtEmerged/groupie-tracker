package mapapi

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func convertCoordinates(s string) (float64, float64, error) {
	coord := strings.Split(s, " ")
	x, err := strconv.ParseFloat(coord[0], 64)
	y, err1 := strconv.ParseFloat(coord[1], 64)
	if err != nil {
		log.Println("3", err)
	}
	if err1 != nil {
		errors.Join(err, err1)
	}
	return x, y, err
}
