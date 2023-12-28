package handler

import (
	"net/http"

	"groupie-tracker/internal/models"
	"groupie-tracker/pkg/groupie"
	mapapi "groupie-tracker/pkg/map"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	artistsPage, relations, err := groupie.Artist()
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}

	id, err := groupie.GetID(artistsPage, r.FormValue("id"))
	if err != nil {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	oneArtist := artistsPage[id]
	oneArtist.Relations = relations.Index[id].DatesLocations
	err = mapapi.RequestMapApi(&oneArtist)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err = models.Tpl.ExecuteTemplate(w, "artist.html", oneArtist)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}
