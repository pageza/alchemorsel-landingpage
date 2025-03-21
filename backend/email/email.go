package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// SendWelcomeEmail sends a welcome email to the provided recipient.
func SendWelcomeEmail(to string) {
	// Load environment variables from .env file (if it exists)
	if err := godotenv.Load(); err != nil {
		log.Println("cursor--No .env file found, using system environment variables")
	}

	// Get environment variables (ensure .env contains these values)
	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")

	// Debug logging (be cautious with logging sensitive information in production)
	log.Printf("cursor--Attempting to send email from %s using SMTP server %s to %s", from, smtpServer, to)

	// SMTP Authentication
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Compose the email message
	subject := "Welcome to AlcheMorsel!"
	body := "Thank you for signing up! We will keep you updated on our launch."

	// Prepare the email headers
	message := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" + // Proper Content-Type
		"\r\n" + // Blank line separates headers from body
		body)

	// Send the email
	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, message)
	if err != nil {
		log.Printf("cursor--Failed to send email: %v", err)
	} else {
		log.Println("cursor--Welcome email sent successfully!")
	}
}
