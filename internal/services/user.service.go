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
	userRepo repositories.UserRepository
}

func NewUserService() UserService {
	return &userService{
		userRepo: repositories.NewUserRepository(),
	}
}

func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.userRepo.Create(user)
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	existingUser, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	existingUser.FirstName = user.FirstName
	existingUser.LastName = user.LastName
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	return s.userRepo.Update(existingUser)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	return s.userRepo.Delete(user)
}
