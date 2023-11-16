package server

import (
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errPage(w, http.StatusNotFound) // 404
		return
	}
	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}
	if index() != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	if initErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err := tpl.ExecuteTemplate(w, "index.html", &IndexPage)
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}