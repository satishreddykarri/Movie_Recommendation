package implementations

import (
	"errors"

	genreDTO "movie-recommendation/internal/dto/genre"
	"movie-recommendation/internal/models"
	repoInterface "movie-recommendation/internal/repositories/interfaces"
	serviceInterface "movie-recommendation/internal/services/interfaces"

	"gorm.io/gorm"
)

type genreService struct {
	repo repoInterface.GenreRepository
}

func NewGenreService(repo repoInterface.GenreRepository) serviceInterface.GenreService {
	return &genreService{
		repo: repo,
	}
}

func (s *genreService) Create(request *genreDTO.CreateGenreRequest) (*genreDTO.GenreResponse, error) {
	_, err := s.repo.FindByName(request.Name)
	if err == nil {
		return nil, errors.New("genre already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	genre := &models.Genre{
		Name: request.Name,
	}
	if err := s.repo.Create(genre); err != nil {
		return nil, err
	}
	return &genreDTO.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}, nil
}

func(s *genreService) GetAll() ([]genreDTO.GenreResponse, error){
	genres, err  := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response :=  make([]genreDTO.GenreResponse, 0, len(genres))
	for _,genre := range genres{
		response = append(response, genreDTO.GenreResponse{
			ID : genre.ID,
			Name : genre.Name,
		})
	}
	return response, nil
}

func (s *genreService) GetByID(id uint) (*genreDTO.GenreResponse, error) {

	genre, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}

		return nil, err
	}

	return &genreDTO.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}, nil
}

func (s *genreService) Update(id uint, request *genreDTO.UpdateGenreRequest) (*genreDTO.GenreResponse, error) {

	genre, err := s.repo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}
		return nil, err
	}

	existingGenre, err := s.repo.FindByName(request.Name)
	if err == nil && existingGenre.ID != genre.ID {
		return nil, errors.New("genre already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	genre.Name = request.Name

	if err := s.repo.Update(genre); err != nil {
		return nil, err
	}

	return &genreDTO.GenreResponse{
		ID:   genre.ID,
		Name: genre.Name,
	}, nil
}

func (s *genreService) Delete(id uint) error{
	_, err := s.repo.FindByID(id)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("genre not found")
		}
		return err
	}
	return s.repo.Delete(id)
}