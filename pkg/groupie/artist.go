package groupie

import (
	"errors"
	"strconv"

	"groupie-tracker/internal/models"
)

func GetID(artistsPage []models.Artists, idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}

	if id < 1 || id > len(artistsPage) || idStr[0] == '0' {
		err = errors.New("not found")
	}
	return id - 1, err
}
