package repository

import (
	"errors"

	"github.com/product/pkg/domain/entity"
	"github.com/product/pkg/adapter/db/model"
	dr "github.com/product/pkg/adapter/db"
	dbConf "github.com/product/internal/config/db"

	"github.com/mitchellh/mapstructure"
	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	repo dr.DbDriver
	db   *gorm.DB

	// data entity.FetchProductsResponse
}

func NewProductRepo(database dbConf.Database) *ProductRepo {
	driver, _ := dr.NewInstanceDb(database)
	return &ProductRepo{
		repo: driver,
		db:   driver.Db().(*gorm.DB),
	}


	// var data entity.FetchProductsResponse
	// products := make([]*entity.Product, 0)
	// products = append(products, &entity.Product{
	// 	Id:		"001",
	// 	Name:	"mouse",
	// 	Price:	200000.00,
	// 	Stock:	10,
	// })
	// products = append(products, &entity.Product{
	// 	Id:		"002",
	// 	Name:	"keyboard",
	// 	Price:	1000000.00,
	// 	Stock:	5,
	// })
	// products = append(products, &entity.Product{
	// 	Id:		"003",
	// 	Name:	"monitor",
	// 	Price:	3000000.00,
	// 	Stock:	3,
	// })

	// data.Products = products

	// return &ProductRepo{
	// 	data: data,
	// }
}

func (pr *ProductRepo) FetchProducts() (interface{}, error) {
	var out entity.FetchProductsResponse
	var res []model.Product
	// var out []*entity.Product
	// out = pr.data

	err := pr.db.Debug().Find(&res).Error
	if err != nil {
		return nil, errors.New("Product(Repo).FetchProducts: Error query : " + err.Error())
	}

	for _, r := range res {
		temp := &entity.Product{
			Id: r.Id,
			Name: r.Name,
			Price: r.Price,
			Stock: r.Stock,
		}
		
		out.Products = append(out.Products, temp)
	}

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
	var res model.Product

	err = pr.db.Debug().Where("id = ?", request.Id).Find(&res).Error

	out.Product = &entity.Product{
		Id: res.Id,
		Name: res.Name,
		Price: res.Price,
		Stock: res.Stock,
	}
	
	// for _, product := range pr.data.Products {
	// 	if request.Id == product.Id {
	// 		out.Product = product
	// 		break
	// 	}
	// }
	
	return out, nil
}