package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func HandleRequest() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Sender data.
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	to := []string{
		"toshvelaga@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(
		"Subject: This is a test email\r\n" + "\r\n" + "Here is a test message\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	emailErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if emailErr != nil {
		fmt.Println(emailErr)
		return
	}
	fmt.Printf("Email successfully sent to %s", to[0])
}

func main() {
	lambda.Start(HandleRequest)
}
