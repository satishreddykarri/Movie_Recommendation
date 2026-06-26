package implementations

import (
	"errors"

	movieDTO "movie-recommendation/internal/dto/movie"
	"movie-recommendation/internal/models"
	repositoryInterface "movie-recommendation/internal/repositories/interfaces"
	serviceInterface "movie-recommendation/internal/services/interfaces"

	"gorm.io/gorm"
)

type movieService struct {
	movieRepo repositoryInterface.MovieRepository
	genreRepo repositoryInterface.GenreRepository
	userRepo repositoryInterface.UserRepository
}

func NewMovieService(
	movieRepo repositoryInterface.MovieRepository,
	genreRepo repositoryInterface.GenreRepository,
	userRepo repositoryInterface.UserRepository,
) serviceInterface.MovieService {

	return &movieService{
		movieRepo: movieRepo,
		genreRepo: genreRepo,
		userRepo: userRepo,
	}
}

func (s *movieService) Create(request *movieDTO.CreateRequest) (*movieDTO.Response, error) {

	_, err := s.genreRepo.FindByID(request.GenreID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}
		return nil, err
	}
	movie := &models.Movie{
		Title:       request.Title,
		Description: request.Description,
		ReleaseYear: request.ReleaseYear,
		GenreID:     request.GenreID,
	}
	if err := s.movieRepo.Create(movie); err != nil {
		return nil, err
	}

	updatedMovie, err := s.movieRepo.FindByID(movie.ID)
	if err != nil {
		return nil, err
	}

	return &movieDTO.Response{
		ID:          updatedMovie.ID,
		Title:       updatedMovie.Title,
		Description: updatedMovie.Description,
		ReleaseYear: updatedMovie.ReleaseYear,
		GenreID:     updatedMovie.GenreID,
		GenreName:   updatedMovie.GenreName,
	}, nil
}

func (s *movieService) GetAll() ([]movieDTO.Response, error) {
	movies, err := s.movieRepo.FindAll()
	if err != nil {
		return nil, err
	}
	response := make([]movieDTO.Response, 0, len(movies))
	for _, movie := range movies {
		response = append(response, movieDTO.Response{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			ReleaseYear: movie.ReleaseYear,
			GenreID:     movie.GenreID,
			GenreName:   movie.GenreName,
		})
	}
	return response, nil
}

func (s *movieService) GetByID(id uint) (*movieDTO.Response, error) {

	movie, err := s.movieRepo.FindModelByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}
		return nil, err
	}

	return &movieDTO.Response{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		ReleaseYear: movie.ReleaseYear,
		GenreID:     movie.GenreID,
	}, nil
}

func (s *movieService) Update(id uint, request *movieDTO.UpdateRequest) (*movieDTO.Response, error) {

	movie, err := s.movieRepo.FindModelByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("movie not found")
		}

		return nil, err
	}

	_, err = s.genreRepo.FindByID(request.GenreID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}

		return nil, err
	}

	movie.Title = request.Title
	movie.Description = request.Description
	movie.ReleaseYear = request.ReleaseYear
	movie.GenreID = request.GenreID

	if err := s.movieRepo.Update(movie); err != nil {
		return nil, err
	}

	updatedMovie, err := s.movieRepo.FindByID(movie.ID)
	if err != nil {
		return nil, err
	}

	return &movieDTO.Response{
		ID:          updatedMovie.ID,
		Title:       updatedMovie.Title,
		Description: updatedMovie.Description,
		ReleaseYear: updatedMovie.ReleaseYear,
		GenreID:     updatedMovie.GenreID,
		GenreName:   updatedMovie.GenreName,
	}, nil
}

func (s *movieService) Delete(id uint) error {

	_, err := s.movieRepo.FindModelByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("movie not found")
		}

		return err
	}

	return s.movieRepo.Delete(id)
}

func (s *movieService) GetByGenreID(genreID uint) ([]movieDTO.Response, error) {

	_, err := s.genreRepo.FindByID(genreID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("genre not found")
		}

		return nil, err
	}

	movies, err := s.movieRepo.FindByGenreID(genreID)
	if err != nil {
		return nil, err
	}

	response := make([]movieDTO.Response, 0, len(movies))

	for _, movie := range movies {

		response = append(response, movieDTO.Response{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			ReleaseYear: movie.ReleaseYear,
			GenreID:     movie.GenreID,
			GenreName:   movie.GenreName,
		})
	}

	return response, nil
}

func (s *movieService) GetMoviesRatedByUser(userID uint) ([]movieDTO.UserMovieResponse, error) {

	_, err := s.userRepo.FindByID(userID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	movies, err := s.movieRepo.GetMoviesRatedByUser(userID)
	if err != nil {
		return nil, err
	}

	response := make([]movieDTO.UserMovieResponse, 0, len(movies))

	for _, movie := range movies {

		response = append(response, movieDTO.UserMovieResponse{
			MovieID:   movie.MovieID,
			MovieName: movie.MovieName,
			Rating:    movie.Rating,
		})
	}

	return response, nil
}

func (s *movieService) GetTopRatedMovies() ([]movieDTO.TopRatedResponse, error) {

	movies, err := s.movieRepo.GetTopRatedMovies()
	if err != nil {
		return nil, err
	}
	response := make([]movieDTO.TopRatedResponse, 0, len(movies))
	for _, movie := range movies {
		response = append(response, movieDTO.TopRatedResponse{
			ID:            movie.ID,
			Title:         movie.Title,
			GenreName:     movie.GenreName,
			AverageRating: movie.AverageRating,
		})
	}
	return response, nil
}

func (s *movieService) GetRecommendedMovies(userID uint) ([]movieDTO.RecommendedMovieResponse, error) {

	_, err := s.userRepo.FindByID(userID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	movies, err := s.movieRepo.GetRecommendedMovies(userID)
	if err != nil {
		return nil, err
	}

	response := make([]movieDTO.RecommendedMovieResponse, 0, len(movies))

	for _, movie := range movies {

		response = append(response, movieDTO.RecommendedMovieResponse{
			ID:            movie.ID,
			Title:         movie.Title,
			GenreName:     movie.GenreName,
			AverageRating: movie.AverageRating,
		})
	}

	return response, nil
}