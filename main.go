package main

// these are the packages we are importing
// if you import it you gotta use it
import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Sender data.
	// Get these env variables from your .env file
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	// replace my email address with yours unless if you want to send me an email :)
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
