package server

import (
	"log"
	"net/http"
)

const (
	address = "localhost"
	port    = "8080"
)

func Running() {
	indexPage, _ := index()
	artistsPage, relation, _ := artist()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, indexPage)
	})
	http.HandleFunc("/artist", func(w http.ResponseWriter, r *http.Request) {
		artistHandler(w, r, artistsPage, relation)
	})
	log.Printf("Listening on: http://%s:%s/\n", address, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
