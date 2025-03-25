package service

import (
	"errors"

	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"
	"AvitoTest/internal/repository"
	"AvitoTest/internal/util"
)

type InfoService struct {
	historyRepo *repository.HistoryRepository
	userRepo    *repository.UserRepository
}

func NewInfoService(
	historyRepo *repository.HistoryRepository,
	userRepo *repository.UserRepository,
) *InfoService {
	return &InfoService{
		historyRepo: historyRepo,
		userRepo:    userRepo,
	}
}

func (s *InfoService) GetUserInfo(userIdStr string) (dto.InfoResponse, error) {
	userId, err := util.ParseUserId(userIdStr)
	if err != nil {
		return dto.InfoResponse{}, err
	}

	infoResponse := dto.InfoResponse{
		Coins:     0,
		Inventory: []dto.Inventory{},
		CoinHistory: dto.CoinHistory{
			Received: []dto.CoinTransaction{},
			Sent:     []dto.CoinTransaction{},
		},
	}

	user, err := s.userRepo.FindAndPreloadUserById(userId)

	if err != nil {
		return infoResponse, errors.New("user exception")
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

	histories, _ := s.historyRepo.FindPreloadHistoryByUserId(user.ID)

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
