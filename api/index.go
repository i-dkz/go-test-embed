package handler

// package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Sender's email address and password
	from := "zachflentgewong@gmail.com"
	appPass := os.Getenv("EMAIL_PASSWORD")

	// Recipient's email address
	to := []string{"zflentge@bestbuycanada.ca"}

	// Gmail SMTP server settings
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	// Parse the JSON request body
	var requestData struct {
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	subject := requestData.Subject
	messageBody := requestData.Body

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
