package movie

type CreateRequest struct {
	Title       string `json:"title" binding:"required,max=255"`
	Description string `json:"description"`
	ReleaseYear int    `json:"release_year" binding:"required"`
	GenreID     uint   `json:"genre_id" binding:"required"`
}