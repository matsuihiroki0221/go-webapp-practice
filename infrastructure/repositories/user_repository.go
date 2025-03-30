package repositories

import (
	"go-webapp-practice/domain/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

type UserEntity struct {
	Id           uint   `gorm:"primaryKey"`
	ProviderId   string `gorm:"uniqueIndex:idx_provider"` // Auth0やGoogleのID
	ProviderName string `gorm:"uniqueIndex:idx_provider"` // "auth0", "google" など
	UserName     string
}

// Entity → Domain（インフラ層で変換）
func (u *UserEntity) ToDomain() *models.User {
	return &models.User{
		Id:           u.Id,
		ProviderId:   u.ProviderId,
		ProviderName: u.UserName,
		UserName:     u.UserName,
	}
}

// Domain → Entity（インフラ層で変換）
func FromDomain(user *models.User) *UserEntity {
	return &UserEntity{
		Id:           user.Id,
		ProviderId:   user.ProviderId,
		ProviderName: user.ProviderName,
		UserName:     user.UserName,
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

func (r *UserRepositoryImpl) FindByProviderID(providerId, providerName string) (*models.User, error) {
	var entity UserEntity
	if err := r.db.Where("provider_id = ? AND provider_name = ?", providerId, providerName).First(&entity).Error; err != nil {
		return nil, err
	}
	return entity.ToDomain(), nil // エンティティ → ドメインモデル変換
}
