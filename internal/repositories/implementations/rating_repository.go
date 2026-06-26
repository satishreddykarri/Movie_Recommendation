package implementations

import (
	"movie-recommendation/internal/models"
	repositoryInterface "movie-recommendation/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) repositoryInterface.RatingRepository {
	return &ratingRepository{
		db: db,
	}
}

func (r *ratingRepository) Create(rating *models.Rating) error {
	return r.db.Create(rating).Error
}

func (r *ratingRepository) FindByUserAndMovie(userID uint, movieID uint) (*models.Rating, error) {

	var rating models.Rating

	err := r.db.
		Where("user_id = ? AND movie_id = ?", userID, movieID).
		First(&rating).Error

	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (r *ratingRepository) FindAll() ([]models.RatingWithDetails, error) {
	var ratings []models.RatingWithDetails
	err := r.db.
		Table("ratings").
		Select(`
			ratings.id,
			ratings.user_id,
			users.name AS user_name,
			ratings.movie_id,
			movies.title AS movie_name,
			ratings.rating
		`).
		Joins("INNER JOIN users ON users.id = ratings.user_id").
		Joins("INNER JOIN movies ON movies.id = ratings.movie_id").
		Order("ratings.id ASC").
		Scan(&ratings).Error
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (r *ratingRepository) FindByID(id uint) (*models.RatingWithDetails, error) {

	var rating models.RatingWithDetails

	err := r.db.
		Table("ratings").
		Select(`
			ratings.id,
			ratings.user_id,
			users.name AS user_name,
			ratings.movie_id,
			movies.title AS movie_name,
			ratings.rating
		`).
		Joins("INNER JOIN users ON users.id = ratings.user_id").
		Joins("INNER JOIN movies ON movies.id = ratings.movie_id").
		Where("ratings.id = ?", id).
		Scan(&rating).Error

	if err != nil {
		return nil, err
	}

	if rating.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &rating, nil
}

func (r *ratingRepository) FindModelByID(id uint) (*models.Rating, error) {

	var rating models.Rating

	err := r.db.First(&rating, id).Error
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func (r *ratingRepository) Update(rating *models.Rating) error {
	return r.db.Save(rating).Error
}

func (r *ratingRepository) Delete(id uint) error {
	return r.db.Delete(&models.Rating{}, id).Error
}