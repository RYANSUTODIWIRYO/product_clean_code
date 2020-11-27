package repository

import (
	"github.com/product/pkg/domain/entity"
)

type ProductRepo struct {
	data entity.FetchProductsResponse
}

func NewProductRepo() *ProductRepo {
	var data entity.FetchProductsResponse
	products := make([]*entity.Product, 0)
	products = append(products, &entity.Product{
		ID:		"001",
		Name:	"mouse",
		Price:	200000.00,
		Stok:	10,
	})
	products = append(products, &entity.Product{
		ID:		"002",
		Name:	"keyboard",
		Price:	1000000.00,
		Stok:	5,
	})
	products = append(products, &entity.Product{
		ID:		"003",
		Name:	"monitor",
		Price:	3000000.00,
		Stok:	3,
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
