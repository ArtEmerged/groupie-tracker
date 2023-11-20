package server

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

func searchPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
		errPage(w, http.StatusNotFound) // 404
		return
	}

	if r.Method != http.MethodGet {
		errPage(w, http.StatusMethodNotAllowed) // 405
		return
	}

	tmp, err := template.ParseFiles("./web/templates/search.html")
	if err != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}

	value := r.FormValue("query")
	list := searchFor(value)

	err = tmp.Execute(w, list)
	if err != nil {
		log.Println("Error:", err)
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
}

func searchFor(value string) []Artists {
	var list []Artists

	for i, v := range Bands {
		
		if strings.Contains(strings.ToLower(v.Name), strings.ToLower(value)) {
			list = append(list, Bands[i])
		}
	}

	return list
}
