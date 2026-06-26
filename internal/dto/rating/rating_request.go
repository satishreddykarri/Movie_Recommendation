package rating

type CreateRequest struct {
	UserID  uint `json:"user_id" binding:"required"`
	MovieID uint `json:"movie_id" binding:"required"`
	Rating  int  `json:"rating" binding:"required,min=1,max=5"`
}