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
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static/"))))
    http.HandleFunc("/", indexPage)
    http.HandleFunc("/artist", artistPage)
    log.Printf("Listening on: http://%s:%s/\n", address, port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}