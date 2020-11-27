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

	// decode response
	// var results *[]entity.Product
	var results *entity.FetchProductsResponse
	err = mapstructure.Decode(res, &results)
	if err != nil {
		return nil, errors.New("Product_Interactor.FetchProducts : Error Decode Response : " + err.Error())
	}

	return pi.out.FetchProductsResponse(results)
}