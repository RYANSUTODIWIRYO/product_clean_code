package usecase

import (
	"errors"

	"github.com/product/pkg/domain/entity"
	"github.com/product/pkg/domain/repository"
	"github.com/mitchellh/mapstructure"
)

type CreateProductInteractor struct {
	repo repository.CreateProductRepository
	out CreateProductOutputPort
}

func NewCreateProductInteractor(db repository.CreateProductRepository, o CreateProductOutputPort) *CreateProductInteractor{
	return &CreateProductInteractor{
		repo: db,
		out: o,
	}
}

func (pi *CreateProductInteractor) CreateProduct(data interface{}) (interface{}, error) {
	// if data is null
	if data == nil {
		return nil, errors.New("Create_Product_Interactor.CreateProduct : Request Body is nil")
	}

	// decode request
	var input *entity.CreateProductRequest
	err := mapstructure.Decode(data, &input)
	if err != nil {
		return nil, errors.New("Create_Product_Interactor.CreateProduct : Failed to Decode Request Body : " + err.Error())
	}
	
	// call product_repository
	res, err := pi.repo.CreateProduct(input)
	if err != nil {
		return nil, errors.New("Create_Product_Interactor.CreateProduct : Failed to Call Repository : " + err.Error())
	}

	// decode response
	var results *entity.CreateProductResponse
	err = mapstructure.Decode(res, &results)
	if err != nil {
		return nil, errors.New("Product_Interactor.FindProductById : Error Decode Response : " + err.Error())
	}

	return pi.out.CreateProductResponse(results)
}