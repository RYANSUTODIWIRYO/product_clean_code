package usecase

import (
	"errors"

	"github.com/product/pkg/domain/entity"
	"github.com/product/pkg/domain/repository"
	"github.com/mitchellh/mapstructure"
)

type ProductInteractor struct {
	repo repository.ProductRepository
	out ProductOutputPort
}

func NewProductInteractor(db repository.ProductRepository, o ProductOutputPort) *ProductInteractor{
	return &ProductInteractor{
		repo: db,
		out: o,
	}
}

func (pi *ProductInteractor) FetchProducts() (interface{}, error) {
	// call product_repository
	res, err := pi.repo.FetchProducts()
	if err != nil {
		return nil, errors.New("Product_Interactor.FetchProducts : Failed to Call Repository : " + err.Error())
	}

	// decode response
	var results *entity.FetchProductsResponse
	err = mapstructure.Decode(res, &results)
	if err != nil {
		return nil, errors.New("Product_Interactor.FetchProducts : Error Decode Response : " + err.Error())
	}

	return pi.out.FetchProductsResponse(results)
}

func (pi *ProductInteractor) FindProductById(data interface{}) (interface{}, error) {
	// if data is null
	if data == nil {
		return nil, errors.New("Product_Interactor.FindProductById : Request Body is nil")
	}

	// decode request
	var input *entity.FindProductByIdRequest
	err := mapstructure.Decode(data, &input)
	if err != nil {
		return nil, errors.New("Product_Interactor.FindProductById : Failed to Decode Request Body : " + err.Error())
	}
	
	// call product_repository
	res, err := pi.repo.FindProductById(input)
	if err != nil {
		return nil, errors.New("Product_Interactor.FindProductById : Failed to Call Repository : " + err.Error())
	}

	// decode response
	var results *entity.FindProductByIdResponse
	err = mapstructure.Decode(res, &results)
	if err != nil {
		return nil, errors.New("Product_Interactor.FindProductById : Error Decode Response : " + err.Error())
	}

	return pi.out.FindProductByIdResponse(results)
}