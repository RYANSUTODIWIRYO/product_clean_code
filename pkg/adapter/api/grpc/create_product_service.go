package grpc

import (
	"context"
	"errors"

	"github.com/product/pkg/domain/entity"
	usecase "github.com/product/pkg/usecase/create_product"
	proto "github.com/product/pkg/infrastructure/grpc/proto/product"

	"github.com/mitchellh/mapstructure"
)

type CreateProductService struct {
	repo usecase.CreateProductInputPort
}

func NewCreateProductService(u usecase.CreateProductInputPort) *CreateProductService {
	return &CreateProductService{
		repo: u,
	}
}

func (s *CreateProductService) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	if req == nil {
		return nil, errors.New("Create_Product_Service.CreateProduct : Request Body is nil")
	}	

	var input *entity.CreateProductRequest
	err := mapstructure.Decode(req, &input)
	if err != nil {
		return nil, errors.New("Create_Product_Service.CreateProduct : Failed to Decode Request Body : " + err.Error())
	}

	// Call Usecase
	res, err := s.repo.CreateProduct(input)
	if err != nil {
		return nil, errors.New("Create_Product_Service.CreateProduct : Failed to Fetch Products : " + err.Error())
	}

	return res.(*proto.CreateProductResponse), nil
}