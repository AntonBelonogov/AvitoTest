package repository

import (
	"AvitoTest/internal/model/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindUserByUsername(username string, tx *gorm.DB) (*entity.User, error)
	CreateUser(user *entity.User) *gorm.DB
	FindAndPreloadUserById(userId uint) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetDB() *gorm.DB {
	return ur.db
}

func (ur *UserRepository) FindUserByUsername(username string, tx *gorm.DB) (*entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User

	result := ur.db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, result.Error
}

func (ur *UserRepository) FindUserById(userId uint, tx *gorm.DB) (*entity.User, error) {
	if tx == nil {
		tx = ur.db
	}

	var user entity.User

	result := ur.db.Where("ID = ?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) && result.Error != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, result.Error
}

func (ur *UserRepository) CreateUser(user *entity.User) (result *gorm.DB) {
	return ur.db.Create(&user)
}

func (ur *UserRepository) FindAndPreloadUserById(userId uint) (*entity.User, error) {
	var user entity.User

	result := ur.db.
		Preload("Product").
		Preload("Product.Product").
		Where("ID = ?", userId).
		First(&user)

	return &user, result.Error
}
