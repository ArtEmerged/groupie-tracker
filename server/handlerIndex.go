package server

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	artistsPage, relations, err := artist()
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	artistsPage = createDateForSearch(artistsPage, relations)
	if initErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err = tpl.ExecuteTemplate(w, "index.html", &artistsPage)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}

func createDateForSearch(artistsPage []artists, relations relation) []artists {
	for i := 0; i < len(artistsPage); i++ {
		artistsPage[i].Relations = relations.Index[i].DatesLocations
	}
	return artistsPage
}
