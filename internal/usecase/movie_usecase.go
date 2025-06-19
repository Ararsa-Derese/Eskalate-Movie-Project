package usecase

import (
	"errors"
	"eskalate-movie-api/internal/domain"
	"eskalate-movie-api/internal/dto"
	"eskalate-movie-api/internal/repository"
	"eskalate-movie-api/pkg/cloudinary"
	"mime/multipart"
	"regexp"

	"github.com/google/uuid"
)

type MovieUsecase struct {
	MovieRepo repository.MovieRepository
}

func NewMovieUsecase(movieRepo repository.MovieRepository) *MovieUsecase {
	return &MovieUsecase{MovieRepo: movieRepo}
}

func (u *MovieUsecase) CreateMovie(req *dto.CreateMovieRequest, posterFile multipart.File, posterHeader *multipart.FileHeader, userID string) (*dto.CreateMovieResponse, error) {
	if !isValidYouTubeURL(req.TrailerUrl) {
		return nil, errors.New("trailerUrl must be a valid YouTube URL")
	}
	posterURL, err := cloudinary.UploadPoster(posterFile, posterHeader)
	if err != nil {
		return nil, errors.New("failed to upload poster")
	}
	movie := &domain.Movie{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		Genres:      req.Genres,
		Actors:      req.Actors,
		Trailer:     req.TrailerUrl,
		Poster:      posterURL,
		UserID:      uuid.MustParse(userID),
	}
	err = u.MovieRepo.Create(movie)
	if err != nil {
		return nil, err
	}
	return &dto.CreateMovieResponse{
		ID:          movie.ID.String(),
		Title:       movie.Title,
		Description: movie.Description,
		Genres:      movie.Genres,
		Actors:      movie.Actors,
		TrailerUrl:  movie.Trailer,
		Poster:      movie.Poster,
	}, nil
}

func isValidYouTubeURL(url string) bool {
	ytRegex := regexp.MustCompile(`^(https?://)?(www\.)?(youtube\.com|youtu\.be)/.+$`)
	return ytRegex.MatchString(url)
}

func (u *MovieUsecase) UpdateMovie(id string, req *dto.UpdateMovieRequest, userID string) (*dto.UpdateMovieResponse, error) {
	movie, err := u.MovieRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if movie.UserID.String() != userID {
		return nil, errors.New("forbidden: you do not own this movie")
	}
	if !isValidYouTubeURL(req.TrailerUrl) {
		return nil, errors.New("trailerUrl must be a valid YouTube URL")
	}
	movie.Title = req.Title
	movie.Description = req.Description
	movie.Genres = req.Genres
	movie.Actors = req.Actors
	movie.Trailer = req.TrailerUrl
	movie.Poster = req.Poster
	if err := u.MovieRepo.Update(movie); err != nil {
		return nil, err
	}
	return &dto.UpdateMovieResponse{
		ID:          movie.ID.String(),
		Title:       movie.Title,
		Description: movie.Description,
		Genres:      movie.Genres,
		Actors:      movie.Actors,
		TrailerUrl:  movie.Trailer,
		Poster:      movie.Poster,
	}, nil
}

func (u *MovieUsecase) GetMovies(req *dto.GetMoviesRequest) (*dto.GetMoviesResponse, error) {
	movies, totalCount, err := u.MovieRepo.GetMovies(req.Page, req.PageSize, req.Title)
	if err != nil {
		return nil, err
	}

	movieResponses := make([]dto.MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = dto.MovieResponse{
			ID:          movie.ID.String(),
			Title:       movie.Title,
			Description: movie.Description,
			Genres:      movie.Genres,
			Actors:      movie.Actors,
			TrailerUrl:  movie.Trailer,
			Poster:      movie.Poster,
		}
	}

	return &dto.GetMoviesResponse{
		Movies:     movieResponses,
		PageNumber: req.Page,
		PageSize:   req.PageSize,
		TotalSize:  totalCount,
	}, nil
}

func (u *MovieUsecase) GetMovieByID(id string) (*dto.MovieDetailsResponse, error) {
	movie, err := u.MovieRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.MovieDetailsResponse{
		ID:          movie.ID.String(),
		Title:       movie.Title,
		Description: movie.Description,
		Genres:      movie.Genres,
		Actors:      movie.Actors,
		TrailerUrl:  movie.Trailer,
		Poster:      movie.Poster,
		UserID:      movie.UserID.String(),
	}, nil
}

func (u *MovieUsecase) DeleteMovie(movieID string, userID string) error {
	// Check if movie exists and belongs to user
	movie, err := u.MovieRepo.FindByID(movieID)
	if err != nil {
		return err
	}

	// Verify ownership
	if movie.UserID.String() != userID {
		return errors.New("forbidden: you do not own this movie")
	}

	// Delete the movie
	return u.MovieRepo.Delete(movieID)
}
