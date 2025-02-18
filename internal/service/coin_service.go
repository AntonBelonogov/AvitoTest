package service

import (
	"errors"
	"strconv"

	"AvitoTest/internal/constants"
	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"
	"AvitoTest/internal/repository"
)

type CoinService struct {
	txRepo      *repository.TransactionRepository
	historyRepo *repository.HistoryRepository
	userRepo    *repository.UserRepository
	productRepo *repository.ProductRepository
}

func NewCoinService(
	txRepo *repository.TransactionRepository,
	historyRepo *repository.HistoryRepository,
	userRepo *repository.UserRepository,
	productRepo *repository.ProductRepository,
) *CoinService {
	return &CoinService{
		txRepo:      txRepo,
		historyRepo: historyRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (cs *CoinService) SendCoin(request dto.SendCoinRequest, userIdStr string) error {
	userId, err := parseUserId(userIdStr)
	if err != nil {
		return err
	}

	requestDto := dto.SendCoinRequestDto{
		FromUserId: userId,
		ToUser:     request.ToUser,
		Amount:     request.Amount,
	}

	tx := cs.txRepo.GetDB().Begin()

	fromUser, err := cs.userRepo.FindUserByIdTx(tx, requestDto.FromUserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if fromUser.Balance < request.Amount {
		tx.Rollback()
		return errors.New("not enough balance")
	}

	toUser, err := cs.userRepo.FindUserByUsernameTx(tx, requestDto.ToUser)
	if err != nil {
		tx.Rollback()
		return err
	}

	fromUser.Balance -= requestDto.Amount
	toUser.Balance += requestDto.Amount

	if err = cs.userRepo.UpdateTx(tx, fromUser); err != nil {
		tx.Rollback()
		return err
	}
	if err = cs.userRepo.UpdateTx(tx, toUser); err != nil {
		tx.Rollback()
		return err
	}

	history := entity.History{
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
		Amount:     requestDto.Amount,
	}

	if err = cs.historyRepo.CreateTx(tx, &history); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (cs *CoinService) BuyItem(item string, userIdStr string) error {
	userId, err := parseUserId(userIdStr)
	if err != nil {
		return err
	}

	tx := cs.txRepo.GetDB().Begin()

	user, err := cs.userRepo.FindUserByIdTx(tx, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	product, err := cs.productRepo.GetProductByNameTx(tx, item)
	if err != nil {
		tx.Rollback()
		return err
	}

	if user.Balance < product.Price {
		tx.Rollback()
		return errors.New("not enough users balance")
	} else {
		user.Balance = user.Balance - product.Price
		if err = cs.userRepo.UpdateTx(tx, user); err != nil {
			tx.Rollback()
			return err
		}
	}

	userProduct := entity.UserProduct{
		UserId:    user.ID,
		ProductId: product.ID,
		Amount:    1,
	}

	if err = cs.productRepo.PutUserProductTx(tx, userProduct); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func parseUserId(userStr string) (uint, error) {
	result, err := strconv.ParseUint(userStr, constants.Base10, constants.BitSize)
	return uint(result), err
}
