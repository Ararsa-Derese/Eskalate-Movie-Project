package handler

import (
	"eskalate-movie-api/internal/dto"
	"eskalate-movie-api/internal/usecase"
	"eskalate-movie-api/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) Signup(c *gin.Context) {
	var req dto.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid input", []string{err.Error()}))
		return
	}

	user, err := h.UserUsecase.Signup(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Signup failed", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusCreated, response.NewSuccessResponse("Signup successful",
		map[string]interface{}{
			"id":       user.ID,
			"email":    user.Email,
			"username": user.Username,
		},
	))
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid input", []string{err.Error()}))
		return
	}

	token, err := h.UserUsecase.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Login failed", []string{err.Error()}))
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessResponse("Login successful",
		dto.LoginResponse{Token: token},
	))
}
