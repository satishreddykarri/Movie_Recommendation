package routes

import (
	"movie-recommendation/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine, handler *handler.MovieHandler) {

	movies := router.Group("/movies")
	{
		movies.POST("", handler.Create)
		movies.GET("", handler.GetAll)
		movies.GET("/top-rated", handler.GetTopRatedMovies)
		movies.GET("/:id", handler.GetByID)
		movies.PUT("/:id", handler.Update)
		movies.DELETE("/:id", handler.Delete)
	}
}