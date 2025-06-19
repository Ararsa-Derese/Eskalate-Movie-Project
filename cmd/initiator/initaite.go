package initiator

import (
	"eskalate-movie-api/internal/domain"
	"eskalate-movie-api/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	// Initialize Gin
	r := gin.Default()

	// Connect to database
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate schema
	dbConn.AutoMigrate(&domain.User{}, &domain.Movie{})

	// Initialize handlers
	handlers := InitializeHandlers(dbConn)

	// Setup routes
	SetupRoutes(r, handlers)

	return r
}
