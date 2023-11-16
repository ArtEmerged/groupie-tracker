package server

import (
	"fmt"
	"net/http"
)

type Err struct {
	Text_err string
	Code_err int
}

func errPage(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	res := Err{Text_err: http.StatusText(code), Code_err: code}
	if initErr != nil {
		errPage(w, http.StatusInternalServerError) // 500
		return
	}
	err := tpl.ExecuteTemplate(w, "error.html", res)
	if err != nil {
		text := fmt.Sprintf("Error 500\n Oppss! %s", http.StatusText(http.StatusInternalServerError))
		http.Error(w, text, http.StatusInternalServerError)
		return
	}
}