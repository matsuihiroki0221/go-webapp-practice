package application

import (
	"go-webapp-practice/domain/models"
)

type UserService struct {
	userRepository UserRepository
}

type UserRepository interface {
	Save(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByProviderID(providerId, providerName string) (*models.User, error)
}

// [アプリケーション層]  →  [ドメイン層]  ←  [インフラ層]
func NewUserService(userRepopository UserRepository) *UserService {
	return &UserService{userRepository: userRepopository}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepository.Save(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepository.FindByID(id)
}
func (s *UserService) GetUserByProviderID(providerId, providerName string) (*models.User, error) {
	return s.userRepository.FindByProviderID(providerId, providerName)
}
