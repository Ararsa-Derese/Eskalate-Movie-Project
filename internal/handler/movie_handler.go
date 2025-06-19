package handler

import (
	"eskalate-movie-api/internal/dto"
	"eskalate-movie-api/internal/usecase"
	"eskalate-movie-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	MovieUsecase *usecase.MovieUsecase
}

func NewMovieHandler(movieUsecase *usecase.MovieUsecase) *MovieHandler {
	return &MovieHandler{MovieUsecase: movieUsecase}
}

func (h *MovieHandler) CreateMovie(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Unauthorized", []string{"unauthorized"}))
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid form data", []string{err.Error()}))
		return
	}

	var req dto.CreateMovieRequest
	req.Title = c.PostForm("title")
	req.Description = c.PostForm("description")
	req.Genres = c.PostFormArray("genres")
	req.Actors = c.PostFormArray("actors")
	req.TrailerUrl = c.PostForm("trailerUrl")

	poster, posterHeader, err := c.Request.FormFile("poster")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Poster is required", []string{err.Error()}))
		return
	}
	defer poster.Close()

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Validation failed", []string{err.Error()}))
		return
	}

	movie, err := h.MovieUsecase.CreateMovie(&req, poster, posterHeader, userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Failed to create movie", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse("Movie created successfully", movie))
}

func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Unauthorized", []string{"unauthorized"}))
		return
	}

	id := c.Param("id")
	var req dto.UpdateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Validation failed", []string{err.Error()}))
		return
	}

	movie, err := h.MovieUsecase.UpdateMovie(id, &req, userID.(string))
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "forbidden: you do not own this movie" {
			status = http.StatusForbidden
		}
		c.JSON(status, response.NewErrorResponse("Failed to update movie", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse("Movie updated successfully", movie))
}

func (h *MovieHandler) GetMovies(c *gin.Context) {
	var req dto.GetMoviesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid pagination parameters", []string{err.Error()}))
		return
	}

	moviesResponse, err := h.MovieUsecase.GetMovies(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to fetch movies", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusOK, response.NewPaginatedResponse(
		"Movies fetched successfully",
		moviesResponse.Movies,
		moviesResponse.PageNumber,
		moviesResponse.PageSize,
		int(moviesResponse.TotalSize),
	))
}

func (h *MovieHandler) GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Movie ID is required", []string{"invalid movie id"}))
		return
	}

	movie, err := h.MovieUsecase.GetMovieByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "movie not found" {
			status = http.StatusNotFound
		}
		c.JSON(status, response.NewErrorResponse("Failed to fetch movie details", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse("Movie details fetched successfully", movie))
}

func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Unauthorized", []string{"unauthorized"}))
		return
	}

	movieID := c.Param("id")
	if movieID == "" {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Movie ID is required", []string{"invalid movie id"}))
		return
	}

	err := h.MovieUsecase.DeleteMovie(movieID, userID.(string))
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "movie not found" {
			status = http.StatusNotFound
		} else if err.Error() == "forbidden: you do not own this movie" {
			status = http.StatusForbidden
		}
		c.JSON(status, response.NewErrorResponse("Failed to delete movie", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse("Movie deleted successfully", nil))
}
