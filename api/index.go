// package handler

package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseFiles("src/templates/index.html"))

func Handler(w http.ResponseWriter, r *http.Request) {

	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("src"))
	router.Handle("GET /src/", http.StripPrefix("/src/", fs))

	router.HandleFunc("GET /", Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
