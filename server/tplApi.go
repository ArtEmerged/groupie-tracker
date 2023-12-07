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
	temp, initErr := template.ParseGlob("./web/templates/*")
	tpl = template.Must(temp, initErr)
}

type artists struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Relations    map[string][]string `json:"-"`
	Markers      []markers
	CountMark    int
}

type relation struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type artistsIndex struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

const (
	artistsAPI  = "https://groupietrackers.herokuapp.com/api/artists"
	relationAPI = "https://groupietrackers.herokuapp.com/api/relation"
)

func index() ([]artistsIndex, error) {
	indexPage := make([]artistsIndex, 0, 50)
	resp, err := http.Get(artistsAPI)
	if err != nil {
		log.Print(err)
		return indexPage, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&indexPage)
	if err != nil {
		log.Print(err)
		return indexPage, err
	}
	return indexPage, err
}

func artist() ([]artists, relation, error) {
	artistsPage := make([]artists, 0, 50)
	relations := relation{}
	resp, err := http.Get(artistsAPI)
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
	resp, err = http.Get(relationAPI)
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
