package initiator

import (
	"eskalate-movie-api/internal/handler"
	"eskalate-movie-api/internal/repository"
	"eskalate-movie-api/internal/usecase"

	"gorm.io/gorm"
)

type Handlers struct {
	UserHandler  *handler.UserHandler
	MovieHandler *handler.MovieHandler
	DocsHandler  *handler.DocsHandler
}

func InitializeHandlers(db *gorm.DB) *Handlers {
	// Initialize repositories
	userRepo := repository.NewPostgresUserRepo(db)
	movieRepo := repository.NewPostgresMovieRepo(db)

	// Initialize use cases
	userUsecase := usecase.NewUserUsecase(userRepo)
	movieUsecase := usecase.NewMovieUsecase(movieRepo)

	// Initialize handlers
	return &Handlers{
		UserHandler:  handler.NewUserHandler(userUsecase),
		MovieHandler: handler.NewMovieHandler(movieUsecase),
		DocsHandler:  handler.NewDocsHandler(),
	}
}
