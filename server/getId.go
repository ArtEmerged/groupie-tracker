package server

import (
	"errors"
	"strconv"
)

func GetID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	if id < 1 || id > len(ArtistsPage) || idStr[0] == '0' {
		err = errors.New("Not Found")
	}
	return id - 1, err
}