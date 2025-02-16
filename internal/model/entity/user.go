package entity

type User struct {
	ID       uint `gorm:"primary_key;AUTO_INCREMENT"`
	Username string
	Password string
	Balance  int
	Product  []UserProduct `gorm:"foreignKey:UserId;references:ID"`
}
