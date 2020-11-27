package usecase

type ProductInputPort interface {
	FetchProducts() (interface{}, error)
}

type ProductOutputPort interface {
	FetchProductsResponse(interface{}) (interface{}, error)
}