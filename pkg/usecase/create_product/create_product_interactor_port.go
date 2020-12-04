package usecase

type CreateProductInputPort interface {
	CreateProduct(interface{}) (interface{}, error)
}

type CreateProductOutputPort interface {
	CreateProductResponse(interface{}) (interface{}, error)
}