package service

import (
	"errors"
	"strconv"

	"AvitoTest/internal/constants"
	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"

	"AvitoTest/internal/repository"
)

type InfoService struct {
	historyRepository repository.HistoryRepositoryInterface
	userRepository    repository.UserRepositoryInterface
}

func NewInfoService(
	infoRepository repository.HistoryRepositoryInterface,
	userRepository repository.UserRepositoryInterface,
) *InfoService {
	return &InfoService{
		historyRepository: infoRepository,
		userRepository:    userRepository,
	}
}

func (s *InfoService) GetUserInfo(userIdStr string) (dto.InfoResponse, error) {
	userId, err := strconv.ParseUint(userIdStr, constants.Base10, constants.BitSize)
	if err != nil {
		return dto.InfoResponse{}, err
	}

	infoResponse := dto.InfoResponse{
		Coins:       constants.ZeroConst,
		Inventory:   []dto.Inventory{},
		CoinHistory: dto.CoinHistory{},
	}

	user, err := s.userRepository.FindAndPreloadUserById(uint(userId))

	if err != nil {
		return infoResponse, errors.New("User exception")
	} else {
		infoResponse.Coins = user.Balance
	}

	var inventoryItems = user.Product

	for _, inventoryItem := range inventoryItems {
		infoResponse.Inventory = append(
			infoResponse.Inventory,
			dto.Inventory{
				Type:     inventoryItem.Product.Name,
				Quantity: inventoryItem.Amount,
			})
	}

	histories, _ := s.historyRepository.FindPreloadHistoryByUserId(user.ID)

	var received []dto.CoinTransaction
	var sent []dto.CoinTransaction

	for _, history := range histories {
		if history.FromUserID == user.ID {
			sent = append(sent, mapToCoinTransaction(history))
		} else {
			received = append(received, mapToCoinTransaction(history))
		}
	}

	infoResponse.CoinHistory.Sent = sent
	infoResponse.CoinHistory.Received = received

	return infoResponse, nil
}

func mapToCoinTransaction(history entity.History) dto.CoinTransaction {
	return dto.CoinTransaction{
		FromUser: history.FromUser.Username,
		ToUser:   history.ToUser.Username,
		Amount:   history.Amount,
	}
}
