package repository

import (
	"gorm.io/gorm"

	"AvitoTest/internal/model/entity"
)

type HistoryRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) *HistoryRepository {
	return &HistoryRepository{db: db}
}

func (hr *HistoryRepository) Create(history *entity.History) error {
	return hr.CreateTx(hr.db, history)
}
func (hr *HistoryRepository) CreateTx(db *gorm.DB, history *entity.History) error {
	return db.Create(&history).Error
}

func (hr *HistoryRepository) FindPreloadHistoryByUserId(userId uint) ([]entity.History, error) {
	var histories []entity.History

	result := hr.db.
		Preload("FromUser").
		Preload("ToUser").
		Where("from_user_id = ? OR to_user_id = ?", userId, userId).
		Find(&histories)

	return histories, result.Error
}

/*func (hr *HistoryRepository) SendCoin(dto dto.SendCoinRequestDto) error {
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
}*/
