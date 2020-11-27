package dto

import (
	"errors"

	proto "github.com/product/pkg/infrastructure/grpc/proto/product"
	"github.com/mitchellh/mapstructure"
)

type ProductBuilder struct {}


func (*ProductBuilder) FetchProductsResponse(data interface{}) (interface{}, error){
	var out *proto.FetchProductsResponse
	err := mapstructure.Decode(data, &out)
	if err != nil {
		return nil, errors.New("ProductBuilder.FetchProductResponse : Failed decode response : " + err.Error())
	}

	return out, nil
}