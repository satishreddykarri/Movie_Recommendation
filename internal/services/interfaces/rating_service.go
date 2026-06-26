package interfaces

import ratingDTO "movie-recommendation/internal/dto/rating"

type RatingService interface {
	Create(request *ratingDTO.CreateRequest) (*ratingDTO.Response, error)
	GetAll() ([]ratingDTO.Response, error)
	GetByID(id uint) (*ratingDTO.Response, error)
	Update(id uint, request *ratingDTO.UpdateRequest) (*ratingDTO.Response, error)
	Delete(id uint) error
}