package repository

import (
	"errors"

	"gorm.io/gorm"

	"AvitoTest/internal/model/entity"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (hr *ProductRepository) Create(product *entity.Product) error {
	return hr.db.Create(&product).Error
}

func (hr *ProductRepository) GetById(id uint) (*entity.Product, error) {
	var product entity.Product

	err := hr.db.Where("id = ?", id).First(&product).Error

	return &product, err
}

func (hr *ProductRepository) Update(product *entity.Product) error {
	if product != nil && product.ID == 0 {
		return errors.New("product ID can't be nil")
	}

	return hr.db.Save(&product).Error
}

func (hr *ProductRepository) Delete(product *entity.Product) error {
	if product != nil && product.ID == 0 {
		return errors.New("product ID can't be nil")
	}

	return hr.db.Delete(&product).Error
}

func (hr *ProductRepository) GetProductByName(productName string) (*entity.Product, error) {
	return hr.GetProductByNameTx(hr.db, productName)
}

func (hr *ProductRepository) GetProductByNameTx(db *gorm.DB, productName string) (*entity.Product, error) {
	var product entity.Product

	err := db.Where("Name = ?", productName).First(&product).Error

	return &product, err
}

func (hr *ProductRepository) SaveUserProductTx(db *gorm.DB, userProduct *entity.UserProduct) error {
	return db.Save(&userProduct).Error
}

func (hr *ProductRepository) PutUserProductTx(tx *gorm.DB, userProduct entity.UserProduct) error {
	if err := tx.
		Where("user_id = ? AND product_id = ?", userProduct.UserId, userProduct.ProductId).
		First(&userProduct).Error; err == nil {
		userProduct.Amount = userProduct.Amount + 1
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := hr.SaveUserProductTx(tx, &userProduct); err != nil {
		return err
	}

	return nil
}
