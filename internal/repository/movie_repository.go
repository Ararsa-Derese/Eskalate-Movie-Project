package repository

import (
	"eskalate-movie-api/internal/domain"

	"errors"

	"gorm.io/gorm"
)

type MovieRepository interface {
	Create(movie *domain.Movie) error
	FindByID(id string) (*domain.Movie, error)
	Update(movie *domain.Movie) error
	GetMovies(page, pageSize int, title string) ([]*domain.Movie, int64, error)
	Delete(id string) error
}

type postgresMovieRepo struct {
	db *gorm.DB
}

func NewPostgresMovieRepo(db *gorm.DB) MovieRepository {
	return &postgresMovieRepo{db: db}
}

func (r *postgresMovieRepo) Create(movie *domain.Movie) error {
	return r.db.Create(movie).Error
}

func (r *postgresMovieRepo) FindByID(id string) (*domain.Movie, error) {
	var movie domain.Movie
	err := r.db.First(&movie, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("movie not found")
	}
	return &movie, err
}

func (r *postgresMovieRepo) Update(movie *domain.Movie) error {
	return r.db.Save(movie).Error
}

func (r *postgresMovieRepo) GetMovies(page, pageSize int, title string) ([]*domain.Movie, int64, error) {
	var movies []*domain.Movie
	var totalCount int64

	// Initialize the query builder
	query := r.db.Model(&domain.Movie{})

	// Add title search condition if title is not empty
	if title != "" {
		query = query.Where("LOWER(title) LIKE LOWER(?)", "%"+title+"%")
	}

	// Get total count with search condition
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated movies with search condition
	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&movies).Error
	if err != nil {
		return nil, 0, err
	}

	return movies, totalCount, nil
}

func (r *postgresMovieRepo) Delete(id string) error {
	result := r.db.Delete(&domain.Movie{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("movie not found")
	}
	return nil
}
