package server

import (
	"fmt"
	"net/http"
)

func artistHandler(w http.ResponseWriter, r *http.Request, artistsPage []artists, relations relation) {
	if r.URL.Path != "/artist" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	if _, _, err := artist(); err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	if initErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	id, err := GetID(artistsPage, r.FormValue("id"))
	if err != nil {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	oneArtist := artistsPage[id]
	oneArtist.Relations = relations.Index[id].DatesLocations
	err = requestMapApi(&oneArtist)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err = tpl.ExecuteTemplate(w, "artist.html", oneArtist)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}
