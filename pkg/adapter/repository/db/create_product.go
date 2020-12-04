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

type CreateProductRepo struct {
	repo dr.DbDriver
	db   *gorm.DB

	// data entity.FetchProductsResponse
}

func NewCreateProductRepo(database dbConf.Database) *CreateProductRepo {
	driver, _ := dr.NewInstanceDb(database)
	return &CreateProductRepo{
		repo: driver,
		db:   driver.Db().(*gorm.DB),
	}
}

func (cpr *CreateProductRepo) CreateProduct(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("Product(Repo).CreateProduct : Request Body is Nil")
	}

	var request *entity.CreateProductRequest
	err := mapstructure.Decode(data, &request)
	if err != nil {
		return nil, errors.New("Product(Repo).CreateProduct : Failed to Decode Request Body : " + err.Error())
	}

	var out entity.CreateProductResponse
	res := model.Product{
		Id:		request.Product.Id,
		Name:	request.Product.Name,
		Price:	request.Product.Price,
		Stock:	request.Product.Stock,
	}

	err = cpr.db.Debug().Create(&res).Error

	if err != nil {
		return nil, errors.New("Product(Repo).CreateProduct: Error query : " + err.Error())
	}

	out.Message = "Insert Product is Succeed"

	return out, nil
}