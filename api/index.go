package handler

// package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Success</h1>")
}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", Handler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
