package service

import (
	"strconv"

	"AvitoTest/internal/constants"
	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/repository"
)

type CoinService struct {
	repository repository.HistoryRepositoryInterface
}

func NewCoinService(repository repository.HistoryRepositoryInterface) *CoinService {
	return &CoinService{repository: repository}
}

func (service *CoinService) SendCoin(request dto.SendCoinRequest, userIdStr string) error {
	userId, err := parseUserId(userIdStr)
	if err != nil {
		return err
	}

	requestDto := dto.SendCoinRequestDto{
		FromUserId: uint(userId),
		ToUser:     request.ToUser,
		Amount:     request.Amount,
	}

	return service.repository.SendCoin(requestDto)
}

func (service *CoinService) BuyItem(item string, userIdStr string) error {
	userId, err := parseUserId(userIdStr)
	if err != nil {
		return err
	}

	return service.repository.BuyItem(item, uint(userId))
}

func parseUserId(userStr string) (uint, error) {
	result, err := strconv.ParseUint(userStr, constants.Base10, constants.BitSize)
	return uint(result), err
}
