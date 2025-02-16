package repository

import (
	"errors"
	"log"

	"gorm.io/gorm"

	"AvitoTest/internal/model/dto"
	"AvitoTest/internal/model/entity"
)

type HistoryRepositoryInterface interface {
	SendCoin(request dto.SendCoinRequestDto) error
	BuyItem(item string, userId uint) error
	FindPreloadHistoryByUserId(userId uint) (histories []entity.History, err error)
}

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (hr *HistoryRepository) FindPreloadHistoryByUserId(userId uint) (histories []entity.History, err error) {
	result := hr.db.
		Preload("FromUser").
		Preload("ToUser").
		Where("from_user_id = ? OR to_user_id = ?", userId, userId).
		Find(&histories)

	return histories, result.Error
}

func (hr *HistoryRepository) SendCoin(dto dto.SendCoinRequestDto) error {
	tx := hr.db.Begin()

	var fromUser entity.User
	if err := tx.Where("ID = ?", dto.FromUserId).First(&fromUser).Error; err != nil {
		tx.Rollback()
	} else if fromUser.Balance < dto.Amount {
		tx.Rollback()
		return errors.New("not enough users balance")
	}

	var toUser entity.User
	if err := tx.Where("username = ?", dto.ToUser).First(&toUser).Error; err != nil {
		tx.Rollback()
	}

	fromUser.Balance = fromUser.Balance - dto.Amount
	toUser.Balance = toUser.Balance + dto.Amount
	tx.Save(fromUser)
	tx.Save(toUser)

	history := entity.History{
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
		Amount:     dto.Amount,
	}

	if err := tx.Create(&history).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (hr *HistoryRepository) BuyItem(item string, userId uint) error {
	tx := hr.db.Begin()

	var user entity.User
	if err := tx.Where("ID = ?", userId).First(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	var product entity.Product
	if err := tx.Where("Name = ?", item).First(&product).Error; err != nil {
		tx.Rollback()
		return errors.New("product not found")
	}

	log.Println("[Buy Item] users + product entity: ", user, product)

	if user.Balance < product.Price {
		tx.Rollback()
		return errors.New("not enough users balance")
	} else {
		user.Balance = user.Balance - product.Price
		tx.Save(&user)
	}

	var userProduct entity.UserProduct

	userProduct = entity.UserProduct{
		UserId:    user.ID,
		ProductId: uint(product.ID),
		Amount:    1,
	}

	if err := tx.
		Where("user_id = ? AND product_id = ?", userId, product.ID).
		First(&userProduct).Error; err != nil {
		tx.Create(&userProduct)
	} else {
		userProduct.Amount = userProduct.Amount + 1
		tx.Save(&userProduct)
	}

	log.Println("[Buy Item] user_product entity: ", user, userProduct)

	return tx.Commit().Error
}
