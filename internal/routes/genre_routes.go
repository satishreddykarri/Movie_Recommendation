package routes

import (
	"movie-recommendation/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterGenreRoutes(router *gin.Engine, handler *handler.GenreHandler,movieHandler *handler.MovieHandler,) {
	genres := router.Group("/genres")
	{
		genres.POST("", handler.Create)
		genres.GET("", handler.GetAll)
		genres.GET("/:id", handler.GetByID)
		genres.PUT("/:id", handler.Update)
		genres.DELETE(":id", handler.Delete)

		genres.GET("/:id/movies", movieHandler.GetByGenreID)
	}
}