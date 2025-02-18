package entity

type History struct {
	ID         uint `gorm:"primary_key;AUTO_INCREMENT"`
	FromUserID uint
	FromUser   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignKey:FromUserID"`
	ToUserID   uint
	ToUser     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignKey:ToUserID"`
	Amount     int
}
