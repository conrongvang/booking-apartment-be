package services

import (
	"booking-apartment/internal/repositories"
	"booking-apartment/models"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.Create(user)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	existingUser, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	return s.repo.Update(existingUser)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
