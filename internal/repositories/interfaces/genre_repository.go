package interfaces
import "movie-recommendation/internal/models"
type GenreRepository interface {
	Create(genre *models.Genre) error
	FindAll() ([]models.Genre, error)
	FindByID(id uint) (*models.Genre, error)
	FindByName(name string) (*models.Genre, error)
	Update(genre *models.Genre) error
	Delete(id uint) error

}
