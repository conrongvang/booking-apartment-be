package repositories

import (
	config "booking-apartment/configs"
	"booking-apartment/models"
	"errors"
)

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("user không tồn tại")
	}
	return &user, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Delete(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
