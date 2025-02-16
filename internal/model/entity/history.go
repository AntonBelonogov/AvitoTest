package entity

type History struct {
	ID         uint `gorm:"primary_key;AUTO_INCREMENT"`
	FromUserID uint
	FromUser   User `gorm:"foreignKey:FromUserID"`
	ToUserID   uint
	ToUser     User `gorm:"foreignKey:ToUserID"`
	Amount     int
}
