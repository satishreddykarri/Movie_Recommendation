package genre
type UpdateGenreRequest struct {
	Name string `json:"name" binding:"required"`
}