package interfaces

import "movie-recommendation/internal/models"

type RatingRepository interface {
	Create(rating *models.Rating) error
	FindByUserAndMovie(userID uint, movieID uint) (*models.Rating, error)
	FindAll() ([]models.RatingWithDetails, error)
	FindByID(id uint) (*models.RatingWithDetails, error)
	FindModelByID(id uint) (*models.Rating, error)
	Update(rating *models.Rating) error
	Delete(id uint) error
}