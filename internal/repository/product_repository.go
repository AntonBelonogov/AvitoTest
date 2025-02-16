package repository

import (
	"gorm.io/gorm"

	"AvitoTest/internal/model/entity"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (hr *ProductRepository) GetProductByNameTx(productName string, tx *gorm.DB) (*entity.Product, error) {
	if tx == nil {
		tx = hr.db
	}

	var product entity.Product

	tx = hr.db.Where("Name = ?", productName).First(&product)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	return &product, tx.Error
}

func (hr *ProductRepository) PutUserProductTx(userId uint, productId uint, tx *gorm.DB) (*entity.UserProduct, error) {
	if tx == nil {
		tx = hr.db
	}

	userProduct := entity.UserProduct{
		UserId:    userId,
		ProductId: productId,
		Amount:    1,
	}

	if err := tx.
		Where("user_id = ? AND product_id = ?", userId, productId).
		First(&userProduct).Error; err != nil {
		tx.Create(&userProduct)
	} else {
		userProduct.Amount = userProduct.Amount + 1
		tx.Save(&userProduct)
	}

	return &userProduct, tx.Error
}
