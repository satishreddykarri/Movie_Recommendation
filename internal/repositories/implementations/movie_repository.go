package implementations

import (
	"movie-recommendation/internal/models"
	repositoryInterface "movie-recommendation/internal/repositories/interfaces"

	"gorm.io/gorm"
)

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) repositoryInterface.MovieRepository {
	return &movieRepository{
		db: db,
	}
}

func (r *movieRepository) Create(movie *models.Movie) error {
	return r.db.Create(movie).Error
}

func (r *movieRepository) FindAll() ([]models.MovieWithGenre, error) {
	var movies []models.MovieWithGenre
	err := r.db.
		Table("movies").
		Select(`
			movies.id,
			movies.title,
			movies.description,
			movies.release_year,
			movies.genre_id,
			genres.name AS genre_name
		`).
		Joins("INNER JOIN genres ON genres.id = movies.genre_id").
		Order("movies.id ASC").
		Scan(&movies).Error
	if err != nil {
		return nil, err
	}
	return movies, nil
}

func (r *movieRepository) FindByID(id uint) (*models.MovieWithGenre, error) {
	var movie models.MovieWithGenre
	err := r.db.
		Table("movies").
		Select(`
			movies.id,
			movies.title,
			movies.description,
			movies.release_year,
			movies.genre_id,
			genres.name AS genre_name
		`).
		Joins("INNER JOIN genres ON genres.id = movies.genre_id").
		Where("movies.id = ?", id).
		Scan(&movie).Error

	if err != nil {
		return nil, err
	}

	if movie.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &movie, nil
}

func (r *movieRepository) FindModelByID(id uint) (*models.Movie, error) {

	var movie models.Movie

	err := r.db.First(&movie, id).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (r *movieRepository) Update(movie *models.Movie) error {
	return r.db.Save(movie).Error
}

func (r *movieRepository) Delete(id uint) error {
	return r.db.Delete(&models.Movie{}, id).Error
}

func (r *movieRepository) FindByGenreID(genreID uint) ([]models.MovieWithGenre, error) {

	var movies []models.MovieWithGenre

	err := r.db.
		Table("movies").
		Select(`
			movies.id,
			movies.title,
			movies.description,
			movies.release_year,
			movies.genre_id,
			genres.name AS genre_name
		`).
		Joins("INNER JOIN genres ON genres.id = movies.genre_id").
		Where("movies.genre_id = ?", genreID).
		Order("movies.id ASC").
		Scan(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *movieRepository) GetMoviesRatedByUser(userID uint) ([]models.UserMovie, error) {

	var movies []models.UserMovie

	err := r.db.
		Table("ratings").
		Select(`
			movies.id AS movie_id,
			movies.title AS movie_name,
			ratings.rating
		`).
		Joins("INNER JOIN movies ON movies.id = ratings.movie_id").
		Where("ratings.user_id = ?", userID).
		Order("movies.title ASC").
		Scan(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *movieRepository) GetTopRatedMovies() ([]models.TopRatedMovie, error) {

	var movies []models.TopRatedMovie

	err := r.db.
		Table("movies").
		Select(`
			movies.id,
			movies.title,
			genres.name AS genre_name,
			AVG(ratings.rating) AS average_rating
		`).
		Joins("INNER JOIN genres ON genres.id = movies.genre_id").
		Joins("INNER JOIN ratings ON ratings.movie_id = movies.id").
		Group(`
			movies.id,
			movies.title,
			genres.name
		`).
		Order("average_rating DESC").
		Scan(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *movieRepository) GetRecommendedMovies(userID uint) ([]models.RecommendedMovie, error) {

	var movies []models.RecommendedMovie

	err := r.db.
		Table("movies").
		Select(`
			movies.id,
			movies.title,
			genres.name AS genre_name,
			AVG(ratings.rating) AS average_rating
		`).
		Joins("INNER JOIN genres ON genres.id = movies.genre_id").
		Joins("INNER JOIN ratings ON ratings.movie_id = movies.id").
		Where(`
			movies.id NOT IN (
				SELECT movie_id
				FROM ratings
				WHERE user_id = ?
			)
		`, userID).
		Group(`
			movies.id,
			movies.title,
			genres.name
		`).
		Order("average_rating DESC").
		Scan(&movies).Error

	if err != nil {
		return nil, err
	}

	return movies, nil
}