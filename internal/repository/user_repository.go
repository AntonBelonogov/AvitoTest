package repository

import (
	"errors"

	"gorm.io/gorm"

	"AvitoTest/internal/model/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(user *entity.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetById(id uint) (*entity.User, error) {
	var user entity.User

	err := ur.db.Where("id = ?", id).First(&user).Error

	return &user, err
}

func (ur *UserRepository) Update(user *entity.User) error {
	return ur.UpdateTx(ur.db, user)
}

func (ur *UserRepository) UpdateTx(tx *gorm.DB, user *entity.User) error {
	if user != nil && user.ID == 0 {
		return errors.New("user ID can't be nil")
	}

	return tx.Save(&user).Error
}

func (ur *UserRepository) Delete(user *entity.User) error {
	if user != nil && user.ID == 0 {
		return errors.New("user ID can't be nil")
	}

	return ur.db.Delete(&user).Error
}

func (ur *UserRepository) FindUserByUsername(username string) (*entity.User, error) {
	return ur.FindUserByUsernameTx(ur.db, username)
}

func (ur *UserRepository) FindUserByUsernameTx(tx *gorm.DB, username string) (*entity.User, error) {
	var user entity.User

	if result := tx.Where("username = ?", username).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) FindUserById(userId uint) (*entity.User, error) {
	return ur.FindUserByIdTx(ur.db, userId)
}

func (ur *UserRepository) FindUserByIdTx(tx *gorm.DB, userId uint) (*entity.User, error) {
	var user entity.User

	if result := tx.Where("ID = ?", userId).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) CreateUser(user *entity.User) (result *gorm.DB) {
	return ur.db.Create(&user)
}

func (ur *UserRepository) FindAndPreloadUserById(userId uint) (*entity.User, error) {
	var user entity.User

	if result := ur.db.
		Preload("Product").
		Preload("Product.Product").
		Where("ID = ?", userId).
		First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
