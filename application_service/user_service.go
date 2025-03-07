package application_service

import (
	"go-webapp-practice/domain/models"
	"go-webapp-practice/infrastructure/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(name, email, password string) error {
	user := &models.User{Name: name, Email: email, Password: password}
	return s.userRepo.Create(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
