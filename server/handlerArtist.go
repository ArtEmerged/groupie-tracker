package server

import (
	"net/http"
)

func artistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artist" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	if artist() != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	if initErr != nil {
		errPage(w, http.StatusInternalServerError)
		return
	}
	id, err := GetID(r.FormValue("id"))
	if err != nil {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	oneArtist := ArtistsPage[id]
	oneArtist.Relations = Relations.Index[id].DatesLocations
	err = tpl.ExecuteTemplate(w, "artist.html", oneArtist)
	if err != nil {
		errPage(w, http.StatusInternalServerError)
		return
	}
}