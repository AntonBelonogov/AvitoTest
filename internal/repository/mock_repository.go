package repository

import (
	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"
	"gorm.io/gorm"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindAndPreloadUserById(userID uint) (*entity.User, error) {
	args := m.Called(userID)
	if user, ok := args.Get(0).(*entity.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) FindUserByUsername(username string, tx *gorm.DB) (*entity.User, error) {
	args := m.Called(username)
	if user, ok := args.Get(0).(*entity.User); ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) CreateUser(user *entity.User) *gorm.DB {
	return m.Called(user).Get(0).(*gorm.DB)
}

type MockHistoryRepository struct {
	mock.Mock
}

func (m *MockHistoryRepository) FindPreloadHistoryByUserId(userID uint) ([]entity.History, error) {
	args := m.Called(userID)
	if histories, ok := args.Get(0).([]entity.History); ok {
		return histories, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockHistoryRepository) SendCoin(request dto.SendCoinRequestDto) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *MockHistoryRepository) BuyItem(item string, userId uint) error {
	args := m.Called(item, userId)
	return args.Error(0)
}
