package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

// SendWelcomeEmail sends a welcome email to the provided recipient.
func SendWelcomeEmail(to string) {
	if err := godotenv.Load(); err != nil {
		log.Println("cursor--No .env file found, using system environment variables")
	}

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")

	// Debug logging (be cautious to not log sensitive info in production)
	log.Printf("cursor--Attempting to send email from %s using SMTP server %s to %s", from, smtpServer, to)

	auth := smtp.PlainAuth("", from, password, smtpServer)

	subject := "Welcome to AlcheMorsel!"
	body := "Thank you for signing up! We will keep you updated on our launch."
	message := []byte("Subject: " + subject + "\r\n" + body)

	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, message)
	if err != nil {
		log.Printf("cursor--Failed to send email: %v", err)
	} else {
		log.Println("cursor--Welcome email sent successfully!")
	}
}
