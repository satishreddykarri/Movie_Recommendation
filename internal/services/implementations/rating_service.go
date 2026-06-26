package implementations

import (
	"errors"

	ratingDTO "movie-recommendation/internal/dto/rating"
	"movie-recommendation/internal/models"
	repositoryInterface "movie-recommendation/internal/repositories/interfaces"
	serviceInterface "movie-recommendation/internal/services/interfaces"

	"gorm.io/gorm"
)

type ratingService struct {
	ratingRepo repositoryInterface.RatingRepository
	userRepo   repositoryInterface.UserRepository
	movieRepo  repositoryInterface.MovieRepository
}

func NewRatingService(
	ratingRepo repositoryInterface.RatingRepository,
	userRepo repositoryInterface.UserRepository,
	movieRepo repositoryInterface.MovieRepository,
) serviceInterface.RatingService {

	return &ratingService{
		ratingRepo: ratingRepo,
		userRepo:   userRepo,
		movieRepo:  movieRepo,
	}
}

func (s *ratingService) Create(request *ratingDTO.CreateRequest) (*ratingDTO.Response, error) {

	user, err := s.userRepo.FindByID(request.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	movie, err := s.movieRepo.FindModelByID(request.MovieID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	_, err = s.ratingRepo.FindByUserAndMovie(request.UserID, request.MovieID)
	if err == nil {
		return nil, errors.New("user has already rated this movie")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	rating := &models.Rating{
		UserID:  request.UserID,
		MovieID: request.MovieID,
		Rating:  request.Rating,
	}

	if err := s.ratingRepo.Create(rating); err != nil {
		return nil, err
	}

	return &ratingDTO.Response{
		ID:        rating.ID,
		UserID:    user.ID,
		UserName:  user.Name,
		MovieID:   movie.ID,
		MovieName: movie.Title,
		Rating:    rating.Rating,
	}, nil
}

func (s *ratingService) GetAll() ([]ratingDTO.Response, error) {

	ratings, err := s.ratingRepo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]ratingDTO.Response, 0, len(ratings))

	for _, rating := range ratings {

		response = append(response, ratingDTO.Response{
			ID:        rating.ID,
			UserID:    rating.UserID,
			UserName:  rating.UserName,
			MovieID:   rating.MovieID,
			MovieName: rating.MovieName,
			Rating:    rating.Rating,
		})
	}

	return response, nil
}

func (s *ratingService) GetByID(id uint) (*ratingDTO.Response, error) {

	rating, err := s.ratingRepo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rating not found")
		}

		return nil, err
	}

	return &ratingDTO.Response{
		ID:        rating.ID,
		UserID:    rating.UserID,
		UserName:  rating.UserName,
		MovieID:   rating.MovieID,
		MovieName: rating.MovieName,
		Rating:    rating.Rating,
	}, nil
}

func (s *ratingService) Update(id uint, request *ratingDTO.UpdateRequest) (*ratingDTO.Response, error) {

	rating, err := s.ratingRepo.FindModelByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rating not found")
		}

		return nil, err
	}

	rating.Rating = request.Rating

	if err := s.ratingRepo.Update(rating); err != nil {
		return nil, err
	}

	updatedRating, err := s.ratingRepo.FindByID(rating.ID)
	if err != nil {
		return nil, err
	}

	return &ratingDTO.Response{
		ID:        updatedRating.ID,
		UserID:    updatedRating.UserID,
		UserName:  updatedRating.UserName,
		MovieID:   updatedRating.MovieID,
		MovieName: updatedRating.MovieName,
		Rating:    updatedRating.Rating,
	}, nil
}

func (s *ratingService) Delete(id uint) error {

	_, err := s.ratingRepo.FindModelByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("rating not found")
		}

		return err
	}

	return s.ratingRepo.Delete(id)
}