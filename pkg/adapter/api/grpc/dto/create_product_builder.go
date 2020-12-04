package dto

import (
	"errors"

	proto "github.com/product/pkg/infrastructure/grpc/proto/product"
	"github.com/mitchellh/mapstructure"
)

type CreateProductBuilder struct {}

func (*CreateProductBuilder) CreateProductResponse(data interface{}) (interface{}, error){
	var out *proto.CreateProductResponse
	err := mapstructure.Decode(data, &out)
	if err != nil {
		return nil, errors.New("CreateProductBuilder.CreateProduct : Failed decode response : " + err.Error())
	}

	return out, nil
}