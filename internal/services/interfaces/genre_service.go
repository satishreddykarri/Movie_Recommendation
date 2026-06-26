package interfaces

import (
	genreDTO "movie-recommendation/internal/dto/genre"
)

type GenreService interface {
	Create(request *genreDTO.CreateGenreRequest) (*genreDTO.GenreResponse, error)
	GetAll() ([]genreDTO.GenreResponse, error)
	GetByID(id uint) (*genreDTO.GenreResponse, error)
	Update(id uint, request *genreDTO.UpdateGenreRequest) (*genreDTO.GenreResponse, error)
	Delete(id uint) error
}