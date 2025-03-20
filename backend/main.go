// cursor--Initial commit: Add Gin API to handle email subscriptions and store them in PostgreSQL.
// cursor--Refactor: Use godotenv to load environment variables from a .env file.
// cursor--Update: Add duplicate email check in /subscribe endpoint.
package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv" // added to load .env file

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file for environment variables; if not found, use system variables
	if err := godotenv.Load(); err != nil {
		log.Println("cursor--No .env file found, using system environment variables")
	}

	// Load environment variables for PostgreSQL credentials
	dbHost := os.Getenv("PG_HOST")
	dbPort := os.Getenv("PG_PORT")
	dbUser := os.Getenv("PG_USER")
	dbPassword := os.Getenv("PG_PASSWORD")
	dbName := os.Getenv("PG_DB")

	// Build PostgreSQL connection string
	psqlInfo := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser +
		" password=" + dbPassword + " dbname=" + dbName + " sslmode=disable"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("cursor--Failed to open database: %v", err)
	}
	defer db.Close()

	// Verify database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("cursor--Failed to ping database: %v", err)
	}
	log.Println("cursor--Connected to database")

	// Initialize Gin router
	router := gin.Default()

	// Enable CORS for frontend interactions
	router.Use(cors.Default())

	// POST /subscribe endpoint to handle email subscriptions
	router.POST("/subscribe", func(c *gin.Context) {
		var req struct {
			Email string `json:"email" binding:"required,email"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email provided."})
			return
		}

		// Check if the email already exists in the database.
		var exists bool
		if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM subscribers WHERE email=$1)", req.Email).Scan(&exists); err != nil {
			log.Printf("cursor--QueryRow error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error."})
			return
		}
		if exists {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists."})
			return
		}

		// Use prepared statement to prevent SQL injection
		stmt, err := db.Prepare("INSERT INTO subscribers (email) VALUES ($1)")
		if err != nil {
			log.Printf("cursor--Prepare error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error."})
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(req.Email)
		if err != nil {
			// Check for duplicate email error (unique constraint violation)
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					c.JSON(http.StatusBadRequest, gin.H{"message": "Email already exists."})
					return
				}
			}
			log.Printf("cursor--Exec error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to subscribe."})
			return
		}
		log.Printf("cursor--New subscriber added: %s", req.Email)
		c.JSON(http.StatusOK, gin.H{"message": "Thanks for signing up!"})
	})

	// cursor--Serve static files from the "public" folder for the frontend.
	router.Static("/", "./public")

	// Start the Gin API on port 8080
	router.Run(":8080")
}
