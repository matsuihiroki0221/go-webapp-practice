package repositories

import (
	"go-webapp-practice/domain/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

type UserEntity struct {
	ID      int `gorm:"primaryKey"`
	Auth0ID string
	Name    string
}

// Entity → Domain（インフラ層で変換）
func (u *UserEntity) ToDomain() *models.User {
	return &models.User{
		ID:      u.ID,
		Auth0ID: u.Auth0ID,
		Name:    u.Name,
	}
}

// Domain → Entity（インフラ層で変換）
func FromDomain(user *models.User) *UserEntity {
	return &UserEntity{
		ID:      user.ID,
		Auth0ID: user.Auth0ID,
		Name:    user.Name,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

// ドメイン層の UserRepository インターフェースを実装
func (r *UserRepositoryImpl) Save(user *models.User) error {
	entity := FromDomain(user) // ドメインモデル → エンティティ変換
	return r.db.Create(entity).Error
}

func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var entity UserEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity.ToDomain(), nil // エンティティ → ドメインモデル変換
}
