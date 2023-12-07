package groupie

import (
	"encoding/json"
	"log"
	"net/http"

	"groupie-tracker/internal/models"
)

func Artist() ([]models.Artists, models.Relation, error) {
	artistsPage := make([]models.Artists, 0, 50)
	relations := models.Relation{}
	resp, err := http.Get(models.ArtistsAPI)
	if err != nil {
		log.Print(err)
		return artistsPage, relations, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artistsPage)
	if err != nil {
		log.Print(err)
		return artistsPage, relations, err
	}
	resp, err = http.Get(models.RelationAPI)
	if err != nil {
		log.Print(err)
		return artistsPage, relations, err
	}
	err = json.NewDecoder(resp.Body).Decode(&relations)
	if err != nil {
		log.Print(err)
		return artistsPage, relations, err
	}
	return artistsPage, relations, err
}
