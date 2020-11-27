package usecase

type ProductInputPort interface {
	FetchProducts() (interface{}, error)
	FindProductById(interface{}) (interface{}, error)
}

type ProductOutputPort interface {
	FetchProductsResponse(interface{}) (interface{}, error)
	FindProductByIdResponse(interface{}) (interface{}, error)
}