package handler

// package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed all:src
var staticFiles embed.FS
var templates = template.Must(template.ParseFS(staticFiles, "src/templates/*.html"))

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		err := templates.ExecuteTemplate(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fs := http.FileServer(http.FS(staticFiles))
		fs.ServeHTTP(w, r)
	}
}

func Main() {
	router := http.NewServeMux()

	// router.Handle("GET /", http.FileServerFS(staticFiles))

	router.HandleFunc("GET /", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
