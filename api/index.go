package handler

// package main

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
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
	} else if r.URL.Path == "/src/style.css" {
		file, err := staticFiles.Open("src/style.css")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer file.Close()

		cssContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/css")
		http.ServeContent(w, r, "style.css", time.Time{}, bytes.NewReader(cssContent))
	} else {
		http.NotFound(w, r)
	}
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
