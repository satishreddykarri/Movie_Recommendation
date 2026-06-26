package routes

import (
	"movie-recommendation/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler, movieHandler *handler.MovieHandler) {

	users := router.Group("/users")
	{
		users.POST("", handler.Create)
		users.GET("", handler.GetAll)
		users.GET("/:id", handler.GetByID)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
		users.GET("/:id/movies", movieHandler.GetMoviesRatedByUser)
		users.GET("/:id/recommendations", movieHandler.GetRecommendedMovies)
	}
}
