package email

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

const welcomeEmailHTML = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Welcome to AlcheMorsel – Your AI Culinary Companion!</title>
  <style>
    body {
      font-family: Arial, sans-serif;
      line-height: 1.6;
      background-color: #f9f9f9;
      color: #333;
      margin: 0;
      padding: 20px;
    }
    .container {
      max-width: 600px;
      background: #fff;
      padding: 20px;
      margin: 0 auto;
      border-radius: 5px;
      box-shadow: 0 0 10px rgba(0,0,0,0.1);
    }
    h1, h2 {
      color: #222;
    }
    ul {
      margin-left: 20px;
    }
    .footer {
      margin-top: 30px;
      font-size: 0.9em;
      color: #777;
    }
  </style>
</head>
<body>
  <div class="container">
    <h1>Welcome to AlcheMorsel – Your AI Culinary Companion!</h1>
    <p>Hi there,</p>
    <p>Thank you for joining <strong>AlcheMorsel</strong> – where we turn everyday ingredients into extraordinary meals!</p>
    <p>At AlcheMorsel, our innovative AI-driven recipe generator takes what you have in your pantry and transforms it into creative, personalized recipes that inspire you to try something new every day.</p>
    <h2>What's in store for you:</h2>
    <ul>
      <li><strong>Smart Recipe Generation:</strong> Simply input your available ingredients, and we'll craft unique recipes tailored just for you.</li>
      <li><strong>Discover New Flavor Combinations:</strong> Break out of your culinary routine with unexpected twists on your favorite dishes.</li>
      <li><strong>Personalized Cooking Experience:</strong> The more you use AlcheMorsel, the better our recommendations become, perfectly matching your tastes.</li>
      <li><strong>Community & Inspiration:</strong> Share your culinary creations and join a community of food enthusiasts eager to explore fresh ideas.</li>
    </ul>
    <p>We're continually refining our platform and adding exciting new features to enhance your cooking journey. Your feedback is invaluable—feel free to reply to this email with any thoughts or suggestions!</p>
    <p>Happy cooking and welcome aboard!</p>
    <p class="footer">Warm regards,<br>The AlcheMorsel Team</p>
  </div>
</body>
</html>
`

// SendWelcomeEmail sends a welcome email to the provided recipient.
func SendWelcomeEmail(to string) {
	if err := godotenv.Load(); err != nil {
		log.Println("cursor--No .env file found, using system environment variables")
	}

	from := os.Getenv("EMAIL_ADDRESS")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")

	// Debug logging (be cautious with logging sensitive information in production)
	log.Printf("cursor--Attempting to send email from %s using SMTP server %s to %s", from, smtpServer, to)

	// SMTP Authentication
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Compose the email message
	subject := "Welcome to AlcheMorsel – Your AI Culinary Companion!"

	// Prepare the email headers and body in HTML format
	message := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" + // Proper Content-Type for HTML
		"\r\n" + // Blank line separates headers from body
		welcomeEmailHTML)

	// Send the email
	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, message)
	if err != nil {
		log.Printf("cursor--Failed to send email: %v", err)
	} else {
		log.Println("cursor--Welcome email sent successfully!")
	}
}
