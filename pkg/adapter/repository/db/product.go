package repository

import (
	"errors"

	"github.com/product/pkg/domain/entity"

	"github.com/mitchellh/mapstructure"
)

type ProductRepo struct {
	data entity.FetchProductsResponse
}

func NewProductRepo() *ProductRepo {
	var data entity.FetchProductsResponse
	products := make([]*entity.Product, 0)
	products = append(products, &entity.Product{
		Id:		"001",
		Name:	"mouse",
		Price:	200000.00,
		Stock:	10,
	})
	products = append(products, &entity.Product{
		Id:		"002",
		Name:	"keyboard",
		Price:	1000000.00,
		Stock:	5,
	})
	products = append(products, &entity.Product{
		Id:		"003",
		Name:	"monitor",
		Price:	3000000.00,
		Stock:	3,
	})

	data.Products = products

	return &ProductRepo{
		data: data,
	}
}

func (pr *ProductRepo) FetchProducts() (interface{}, error) {
	var out entity.FetchProductsResponse
	// var out []*entity.Product
	out = pr.data

	return out, nil
}

func (pr *ProductRepo) FindProductById(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("Product(Repo).FindProductById : Request Body is Nil")
	}

	var request *entity.FindProductByIdRequest
	err := mapstructure.Decode(data, &request)
	if err != nil {
		return nil, errors.New("Product(Repo).FindProductById : Failed to Decode Request Body : " + err.Error())
	}

	var out entity.FindProductByIdResponse
	
	for _, product := range pr.data.Products {
		if request.Id == product.Id {
			out.Product = product
			break
		}
	}
	
	return out, nil
}