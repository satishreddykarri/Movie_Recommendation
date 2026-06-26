package handler

import (
	genreDTO "movie-recommendation/internal/dto/genre"
	serviceInterface "movie-recommendation/internal/services/interfaces"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service serviceInterface.GenreService
}

func NewGenreHandler(service serviceInterface.GenreService) *GenreHandler {
	return &GenreHandler{
		service: service,
	}
}

func (h *GenreHandler) Create(c *gin.Context) {
	var request genreDTO.CreateGenreRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := h.service.Create(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (h *GenreHandler) GetAll(c *gin.Context){
	response, err := h.service.GetAll()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *GenreHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid genre id",
		})
		return
	}

	response, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *GenreHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid genre id",
		})
		return
	}
	var request genreDTO.UpdateGenreRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := h.service.Update(uint(id), &request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *GenreHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid genre id",
		})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "genre deleted successfully",
	})
}
