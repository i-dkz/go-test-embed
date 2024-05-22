package handler

// package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Init(w http.ResponseWriter) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		log.Fatal("Error loading .env file")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Sender's email address and password
	from := "zachflentgewong@gmail.com"
	appPass := os.Getenv("EMAIL_PASSWORD")

	// Recipient's email address
	to := []string{"zachflentgewong@gmail.com"}

	// Gmail SMTP server settings
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the subject and body from the request body
	requestBody := string(body)
	fmt.Println("Request Body:", requestBody)

	parts := strings.Split(requestBody, "&")
	if len(parts) != 2 {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var subject, messageBody string
	for _, part := range parts {
		if strings.HasPrefix(part, "subject=") {
			subject = strings.TrimPrefix(part, "subject=")
		} else if strings.HasPrefix(part, "body=") {
			messageBody = strings.TrimPrefix(part, "body=")
		}
	}

	fmt.Println("Subject:", subject)
	fmt.Println("Message Body:", messageBody)

	// Compose the email message
	message := []byte("Subject: " + subject + "\n\n" + messageBody)

	// Authentication
	auth := smtp.PlainAuth("", from, appPass, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	fmt.Println("Email sent successfully!")
	fmt.Sprintln("Sending email")
	w.WriteHeader(http.StatusOK)

}

func Main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /", Handler)

	log.Println("LISTENING AT PORT:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
