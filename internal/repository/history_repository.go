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
