package server

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

func GetID(artistsPage []artists, idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	if id < 1 || id > len(artistsPage) || idStr[0] == '0' {
		err = errors.New("not found")
	}
	return id - 1, err
}

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
