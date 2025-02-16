package service

import (
	"errors"
	"testing"

	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestCoinService_SendCoin(t *testing.T) {
	mockRepo := new(repository.MockHistoryRepository)
	service := NewCoinService(mockRepo)

	request := dto.SendCoinRequest{
		ToUser: "testUser",
		Amount: 500,
	}

	userIdStr := "3"
	userId := uint(3)

	expectedRequest := dto.SendCoinRequestDto{
		FromUserId: userId,
		ToUser:     "testUser",
		Amount:     500,
	}

	t.Run("Успешная отправка монет", func(t *testing.T) {
		mockRepo.On("SendCoin", expectedRequest).Return(nil).Once()

		err := service.SendCoin(request, userIdStr)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Ошибка: некорректный userId", func(t *testing.T) {
		err := service.SendCoin(request, "invalid_id")

		assert.Error(t, err)
	})

	t.Run("Ошибка при отправке монет", func(t *testing.T) {
		mockRepo.On("SendCoin", expectedRequest).Return(errors.New("DB error")).Once()

		err := service.SendCoin(request, userIdStr)

		assert.Error(t, err)
		assert.Equal(t, "DB error", err.Error())

		mockRepo.AssertExpectations(t)
	})
}

func TestCoinService_BuyItem(t *testing.T) {
	mockRepo := new(repository.MockHistoryRepository)
	service := NewCoinService(mockRepo)

	item := "Sword"
	userIdStr := "5"
	userId := uint(5)

	t.Run("Успешная покупка предмета", func(t *testing.T) {
		mockRepo.On("BuyItem", item, userId).Return(nil).Once()

		err := service.BuyItem(item, userIdStr)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Ошибка: некорректный userId", func(t *testing.T) {
		err := service.BuyItem(item, "invalid_id")

		assert.Error(t, err)
	})

	t.Run("Ошибка при покупке предмета", func(t *testing.T) {
		mockRepo.On("BuyItem", item, userId).Return(errors.New("DB error")).Once()

		err := service.BuyItem(item, userIdStr)

		assert.Error(t, err)
		assert.Equal(t, "DB error", err.Error())

		mockRepo.AssertExpectations(t)
	})
}
