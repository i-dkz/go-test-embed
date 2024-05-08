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

	router.Handle("GET /src/", http.FileServerFS(staticFiles))

	router.HandleFunc("/", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
