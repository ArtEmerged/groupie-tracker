package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var (
	tpl     *template.Template
	initErr error
)

func init() {
	temp, initErr := template.ParseGlob("./web/templates/*.html")
	tpl = template.Must(temp, initErr)
}

type Artists struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"-"`
}

type Relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ArtistsIndex struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

const (
	artistsAPI  = "https://groupietrackers.herokuapp.com/api/artists"
	relationAPI = "https://groupietrackers.herokuapp.com/api/relation"
)

var (
	IndexPage   []ArtistsIndex
	ArtistsPage []Artists
	Relations   Relation
)

func index() error {
	resp, err := http.Get(artistsAPI)
	if err != nil {
		log.Print(err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&IndexPage)
	if err != nil {
		log.Print(err)
		return err
	}
	return err
}

func artist() error {
	resp, err := http.Get(artistsAPI)
	if err != nil {
		log.Print(err)
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&ArtistsPage)
	if err != nil {
		log.Print(err)
		return err
	}
	resp, err = http.Get(relationAPI)
	if err != nil {
		log.Print(err)
		return err
	}
	err = json.NewDecoder(resp.Body).Decode(&Relations)
	if err != nil {
		log.Print(err)
		return err
	}
	return err
}