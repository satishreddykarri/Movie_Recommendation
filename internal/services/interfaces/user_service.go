package interfaces

import userDTO "movie-recommendation/internal/dto/user"

type UserService interface {
	Create(request *userDTO.CreateRequest) (*userDTO.Response, error)
	GetAll() ([]userDTO.Response, error)
	GetByID(id uint) (*userDTO.Response, error)
	Update(id uint, request *userDTO.UpdateRequest) (*userDTO.Response, error)
	Delete(id uint) error
}