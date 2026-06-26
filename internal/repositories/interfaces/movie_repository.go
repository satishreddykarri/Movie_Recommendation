package interfaces

import (
	"movie-recommendation/internal/models"
)

type MovieRepository interface {
	Create(movie *models.Movie) error
	FindAll() ([]models.MovieWithGenre, error)
	FindByID(id uint) (*models.MovieWithGenre, error)
	FindModelByID(id uint) (*models.Movie, error)
	Update(movie *models.Movie) error
	Delete(id uint) error
	FindByGenreID(genreID uint) ([]models.MovieWithGenre, error)
	GetMoviesRatedByUser(userID uint) ([]models.UserMovie, error)
	GetTopRatedMovies() ([]models.TopRatedMovie, error)
	GetRecommendedMovies(userID uint) ([]models.RecommendedMovie, error)
}