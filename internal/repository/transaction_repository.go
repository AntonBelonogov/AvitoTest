package repository

import (
	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (hr *TransactionRepository) GetDB() *gorm.DB {
	return hr.db
}
