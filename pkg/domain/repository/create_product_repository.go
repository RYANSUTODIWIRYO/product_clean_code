package repository

type CreateProductRepository interface {
	CreateProduct(interface{}) (interface{}, error)
}