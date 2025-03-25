package entity

type Product struct {
	ID    uint   `json:"-" gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
