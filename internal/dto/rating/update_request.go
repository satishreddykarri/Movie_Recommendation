package rating

type UpdateRequest struct {
	Rating int `json:"rating" binding:"required,min=1,max=5"`
}