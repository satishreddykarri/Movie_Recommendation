package user

type UpdateRequest struct {
	Name  string `json:"name" binding:"required,max=100"`
	Email string `json:"email" binding:"required,email,max=255"`
}