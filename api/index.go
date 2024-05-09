package handler

// package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed home.html style.css
var staticFiles embed.FS
var templates = template.Must(template.ParseFS(staticFiles, "home.html"))

func Handler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "home.html", nil)

	fmt.Fprintf(w, "<h1>Route: %s</h1>", r.URL.Path)
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
