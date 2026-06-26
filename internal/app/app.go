package app

import (
	"movie-recommendation/internal/handler"
	"movie-recommendation/internal/repositories/implementations"
	"movie-recommendation/internal/routes"
	serviceImpl "movie-recommendation/internal/services/implementations"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	// Repositories
	genreRepo := implementations.NewGenreRepository(db)
	movieRepo := implementations.NewMovieRepository(db)
	userRepo := implementations.NewUserRepository(db)
	ratingRepo := implementations.NewRatingRepository(db)

	// Services
	genreService := serviceImpl.NewGenreService(genreRepo)

	movieService := serviceImpl.NewMovieService(
		movieRepo,
		genreRepo,
		userRepo,
	)

	userService := serviceImpl.NewUserService(userRepo)

	ratingService := serviceImpl.NewRatingService(
		ratingRepo,
		userRepo,
		movieRepo,
	)

	// Handlers
	genreHandler := handler.NewGenreHandler(genreService)
	movieHandler := handler.NewMovieHandler(movieService)
	userHandler := handler.NewUserHandler(userService)
	ratingHandler := handler.NewRatingHandler(ratingService)

	// Routes
	routes.RegisterGenreRoutes(
		router,
		genreHandler,
		movieHandler,
	)

	routes.RegisterMovieRoutes(router, movieHandler)

	routes.RegisterUserRoutes(
		router,
		userHandler,
		movieHandler,
	)

	routes.RegisterRatingRoutes(router, ratingHandler)

	return router
}