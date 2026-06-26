package implementations

import(
	"movie-recommendation/internal/models"
	repoInterface "movie-recommendation/internal/repositories/interfaces"
	"gorm.io/gorm"
)

type genreRepository struct{
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) repoInterface.GenreRepository {
	return &genreRepository{
		db: db,
	}
}

func (r *genreRepository) Create(genre *models.Genre) error {
	return r.db.Create(genre).Error
}

func (r *genreRepository) FindByName(name string) (*models.Genre, error) {
	var genre models.Genre
	err := r.db.
		Where("name = ?", name).
		First(&genre).Error
	if err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *genreRepository) FindAll() ([]models.Genre, error) {
	var genres []models.Genre

	err := r.db.
		Order("name ASC").
		Find(&genres).Error

	if err != nil {
		return nil, err
	}

	return genres, nil
}

func (r *genreRepository) FindByID(id uint) (*models.Genre, error) {
	var genre models.Genre

	err := r.db.
		Where("id = ?", id).
		First(&genre).Error

	if err != nil {
		return nil, err
	}

	return &genre, nil
}

func (r *genreRepository) Update(genre *models.Genre) error {
	return r.db.Save(genre).Error
}

func (r *genreRepository) Delete(id uint) error{
	return r.db.Delete(&models.Genre{}, id).Error
}