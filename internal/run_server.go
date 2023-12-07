package internal

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/internal/handler"
	"groupie-tracker/internal/models"
)

func init() {
	temp, initErr := template.ParseGlob("./web/templates/*")
	models.Tpl = template.Must(temp, initErr)
}

func Running() {
	fileServer := http.FileServer(http.Dir("./web/static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.HandleFunc("/filters/", handler.FilterHandler)
	log.Printf("Listening on: http://%s:%s/\n", models.Address, models.Port)
	err := http.ListenAndServe(":"+models.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
