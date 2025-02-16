package entity

type Product struct {
	ID    int64  `json:"-" gorm:"primary_key;AUTO_INCREMENT"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
