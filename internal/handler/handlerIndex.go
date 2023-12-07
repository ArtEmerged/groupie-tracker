package handler

import (
	"net/http"

	"groupie-tracker/internal/models"
	"groupie-tracker/pkg/groupie"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
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
	artistsPage = createDateForSearch(artistsPage, relations)
	if models.InitErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err = models.Tpl.ExecuteTemplate(w, "index.html", &artistsPage)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}

func createDateForSearch(artistsPage []models.Artists, relations models.Relation) []models.Artists {
	for i := 0; i < len(artistsPage); i++ {
		artistsPage[i].Relations = relations.Index[i].DatesLocations
	}
	return artistsPage
}
