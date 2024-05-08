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
	templates.ExecuteTemplate(w, "index.html", nil)
}

func Main() {
	router := http.NewServeMux()

	fs := http.FS(staticFiles)
	router.Handle("GET /src/style.css", http.FileServer(fs))

	router.HandleFunc("/", Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
