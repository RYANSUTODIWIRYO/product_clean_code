package repository

type ProductRepository interface {
	FetchProducts() (interface{}, error)
	FindProductById(interface{}) (interface{}, error)
}