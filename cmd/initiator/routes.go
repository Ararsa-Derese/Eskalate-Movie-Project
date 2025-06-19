package initiator

import (
	"eskalate-movie-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h *Handlers) {
	// Apply global middleware
	r.Use(middleware.ErrorHandler())

	// Documentation routes
	r.GET("/docs", h.DocsHandler.ServeSwaggerUI)
	r.GET("/swagger.yaml", h.DocsHandler.ServeSwaggerYAML)

	// Auth routes
	auth := r.Group("/")
	{
		auth.POST("/signup", h.UserHandler.Signup)
		auth.POST("/login", h.UserHandler.Login)
	}

	// Movie routes
	movies := r.Group("/movies")
	{
		// Public routes
		movies.GET("", h.MovieHandler.GetMovies)
		movies.GET("/:id", h.MovieHandler.GetMovieByID)

		// Protected routes
		protected := movies.Use(middleware.AuthMiddleware())
		{
			protected.POST("", h.MovieHandler.CreateMovie)
			protected.PUT("/:id", h.MovieHandler.UpdateMovie)
			protected.DELETE("/:id", h.MovieHandler.DeleteMovie)
		}
	}
}
