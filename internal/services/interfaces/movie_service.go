package interfaces

import (
	movieDTO "movie-recommendation/internal/dto/movie"
)

type MovieService interface {
	Create(request *movieDTO.CreateRequest) (*movieDTO.Response, error)
	GetAll() ([]movieDTO.Response, error)
	GetByID(id uint) (*movieDTO.Response, error)
	Update(id uint, request *movieDTO.UpdateRequest) (*movieDTO.Response, error)
	Delete(id uint) error
	GetByGenreID(genreID uint) ([]movieDTO.Response, error)
	GetMoviesRatedByUser(userID uint) ([]movieDTO.UserMovieResponse, error)
	GetTopRatedMovies() ([]movieDTO.TopRatedResponse, error)
	GetRecommendedMovies(userID uint) ([]movieDTO.RecommendedMovieResponse, error)
}