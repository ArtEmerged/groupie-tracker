package handler

import (
	"net/http"

	"groupie-tracker/internal/models"
	"groupie-tracker/pkg/filter"
	"groupie-tracker/pkg/groupie"
)

func FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filters/" {
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

	confing, err := filter.FilterConfig(artistsPage)
	if err != nil {
		errPage(w, http.StatusBadRequest) // 400
		return
	}

	artistsPage = createDateForSearch(artistsPage, relations)
	list, err := filter.FilterFor(r, artistsPage, confing)
	if err != nil {
		errPage(w, http.StatusBadRequest) // 400
		return
	}
	err = models.Tpl.ExecuteTemplate(w, "index.html", list)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}
