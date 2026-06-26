package routes

import (
	"movie-recommendation/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRatingRoutes(router *gin.Engine, handler *handler.RatingHandler) {

	ratings := router.Group("/ratings")
	{
		ratings.POST("", handler.Create)
		ratings.GET("", handler.GetAll)
		ratings.GET("/:id", handler.GetByID)
		ratings.PUT("/:id", handler.Update)
		ratings.DELETE("/:id", handler.Delete)
	}
}
