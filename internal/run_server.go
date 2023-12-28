package internal

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/internal/handler"
	"groupie-tracker/internal/models"
)

func init() {
	temp, err := template.ParseGlob("./web/templates/*.html")
	if err != nil {
		log.Fatalln(err)
	}
	models.Tpl = temp
}

func Running(port string) {
	fileServer := http.FileServer(http.Dir("./web/static/"))

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/artist", handler.ArtistHandler)
	http.HandleFunc("/filters/", handler.FilterHandler)

	log.Printf("Listening on: http://localhost:%s/\n", port)
	
	err := http.ListenAndServe(":"+port, nil)


	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
