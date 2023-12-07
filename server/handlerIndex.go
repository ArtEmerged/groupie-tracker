package server

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request, indexPage []artistsIndex) {
	if r.URL.Path != "/" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	if _, err := index(); err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	if initErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err := tpl.ExecuteTemplate(w, "index.html", &indexPage)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}
