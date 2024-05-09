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

	fmt.Fprintf(w, `
	<!DOCTYPE html>
		<html lang="en">
  			<head>
    			<meta charset="UTF-8" />
    			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
    			<title>Success</title>
    			<link rel="stylesheet" href="/api/style.css" />
  			</head>
  			<body>
    			<h1>Route: %s
    			</h1>
  			</body>
		</html>
	`, r.URL.Path)
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)
	router.HandleFunc("GET https://go-test-embed.vercel.app/style.css", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
