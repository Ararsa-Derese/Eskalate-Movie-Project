package usecase

import (
	"errors"
	"eskalate-movie-api/internal/domain"
	"eskalate-movie-api/internal/dto"
	"eskalate-movie-api/internal/repository"
	"eskalate-movie-api/pkg/security"
	"regexp"

	"github.com/google/uuid"
)

type UserUsecase struct {
	UserRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo}
}

func (u *UserUsecase) Signup(req *dto.SignupRequest) (*domain.User, error) {
	// Email and username uniqueness checked in repo
	if !isValidPassword(req.Password) {
		return nil, errors.New("password must be at least 8 characters, include special, uppercase, and lowercase characters")
	}

	hash, err := security.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &domain.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Username: req.Username,
		Password: hash,
	}

	err = u.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasSpecial = regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString
	)
	return hasUpper(password) && hasLower(password) && hasSpecial(password)
}

func (u *UserUsecase) Login(req *dto.LoginRequest) (string, error) {
	user, err := u.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}
	if !security.CheckPasswordHash(req.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}
	token, err := security.GenerateJWT(user.ID.String(), user.Email)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
