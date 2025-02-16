package service

import (
	"AvitoTest/internal/repository"
	"errors"
	"testing"

	"AvitoTest/internal/model/entity"
	"github.com/stretchr/testify/assert"
)

func TestInfoService_GetUserInfo(t *testing.T) {
	mockUserRepo := new(repository.MockUserRepository)
	mockHistoryRepo := new(repository.MockHistoryRepository)
	service := NewInfoService(mockHistoryRepo, mockUserRepo)

	t.Run("Успешное получение информации о пользователе", func(t *testing.T) {
		userID := uint(1)
		user := &entity.User{
			ID:       userID,
			Username: "testUser",
			Balance:  1000,
			Product: []entity.UserProduct{
				{Product: entity.Product{Name: "Item1"}, Amount: 3},
				{Product: entity.Product{Name: "Item2"}, Amount: 5},
			},
		}

		histories := []entity.History{
			{FromUserID: 1, ToUser: entity.User{Username: "Receiver1"}, Amount: 100},
			{FromUserID: 2, FromUser: entity.User{Username: "Sender1"}, ToUserID: 1, Amount: 50},
		}

		mockUserRepo.On("FindAndPreloadUserById", userID).Return(user, nil)
		mockHistoryRepo.On("FindPreloadHistoryByUserId", userID).Return(histories, nil)

		response, err := service.GetUserInfo("1")

		assert.NoError(t, err)
		assert.EqualValues(t, 1000, response.Coins)
		assert.Len(t, response.Inventory, 2)
		assert.Len(t, response.CoinHistory.Sent, 1)
		assert.Len(t, response.CoinHistory.Received, 1)

		mockUserRepo.AssertExpectations(t)
		mockHistoryRepo.AssertExpectations(t)
	})

	t.Run("Ошибка при конвертации userIdStr", func(t *testing.T) {
		_, err := service.GetUserInfo("invalidID")
		assert.Error(t, err)
	})

	t.Run("Ошибка при поиске пользователя", func(t *testing.T) {
		mockUserRepo.On("FindAndPreloadUserById", uint(2)).Return((*entity.User)(nil), errors.New("User not found"))

		_, err := service.GetUserInfo("2")
		assert.Error(t, err)

		mockUserRepo.AssertExpectations(t)
	})

	t.Run("Ошибка при загрузке истории транзакций", func(t *testing.T) {
		userID := uint(3)
		user := &entity.User{
			ID:       userID,
			Username: "testUser",
			Balance:  1000,
			Product:  []entity.UserProduct{},
		}

		mockUserRepo.On("FindAndPreloadUserById", userID).Return(user, nil)
		mockHistoryRepo.On("FindPreloadHistoryByUserId", userID).Return(nil, errors.New("DB error"))

		response, err := service.GetUserInfo("3")

		assert.NoError(t, err)
		assert.EqualValues(t, 1000, response.Coins)
		assert.Len(t, response.CoinHistory.Sent, 0)
		assert.Len(t, response.CoinHistory.Received, 0)

		mockUserRepo.AssertExpectations(t)
		mockHistoryRepo.AssertExpectations(t)
	})
}
