package implementations

import (
	"errors"

	userDTO "movie-recommendation/internal/dto/user"
	"movie-recommendation/internal/models"
	repositoryInterface "movie-recommendation/internal/repositories/interfaces"
	serviceInterface "movie-recommendation/internal/services/interfaces"

	"gorm.io/gorm"
)

type userService struct {
	repo repositoryInterface.UserRepository
}

func NewUserService(repo repositoryInterface.UserRepository) serviceInterface.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(request *userDTO.CreateRequest) (*userDTO.Response, error) {

	_, err := s.repo.FindByEmail(request.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user := &models.User{
		Name:  request.Name,
		Email: request.Email,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &userDTO.Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) GetAll() ([]userDTO.Response, error) {

	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]userDTO.Response, 0, len(users))

	for _, user := range users {

		response = append(response, userDTO.Response{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	return response, nil
}

func (s *userService) GetByID(id uint) (*userDTO.Response, error) {

	user, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	return &userDTO.Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) Update(id uint, request *userDTO.UpdateRequest) (*userDTO.Response, error) {

	user, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, err
	}

	existingUser, err := s.repo.FindByEmail(request.Email)
	if err == nil && existingUser.ID != user.ID {
		return nil, errors.New("email already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	user.Name = request.Name
	user.Email = request.Email

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	return &userDTO.Response{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) Delete(id uint) error {

	_, err := s.repo.FindByID(id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}

		return err
	}

	return s.repo.Delete(id)
}