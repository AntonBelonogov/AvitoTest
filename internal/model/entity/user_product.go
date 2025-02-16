package entity

type UserProduct struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignKey:ID;"`
	ProductId uint
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL,foreignKey:ID;"`
	Amount    int
}
