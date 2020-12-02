package model

type Product struct {
	Id		string `gorm:"column:id"`
	Name	string `gorm:"column:name"`
	Price	float64 `gorm:"column:price"`
	Stock	int `gorm:"column:stock"`
}

func (Product) TableName() string {
	return "product"
}