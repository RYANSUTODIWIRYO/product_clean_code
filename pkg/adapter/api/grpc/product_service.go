package grpc

import (
	"context"
	"errors"

	"github.com/product/pkg/domain/entity"
	usecase "github.com/product/pkg/usecase/product"
	proto "github.com/product/pkg/infrastructure/grpc/proto/product"

	"github.com/mitchellh/mapstructure"
)

type ProductService struct {
	repo usecase.ProductInputPort
}

func NewProductService(u usecase.ProductInputPort) *ProductService {
	return &ProductService{
		repo: u,
	}
}

func (s *ProductService) FetchProducts(ctx context.Context, req *proto.Empty) (*proto.FetchProductsResponse, error) {
	// Call Usecase
	res, err := s.repo.FetchProducts()
	if err != nil {
		return nil, errors.New("Product_Service.FetchProducts : Failed to Fetch Products : " + err.Error())
	}

	return res.(*proto.FetchProductsResponse), nil
}

func (s *ProductService) FindProductById(ctx context.Context, req *proto.FindProductByIdRequest) (*proto.FindProductByIdResponse, error) {
	if req == nil {
		return nil, errors.New("Product_Service.FindProductById : Request Body is nil")
	}	

	var input *entity.FindProductByIdRequest
	err := mapstructure.Decode(req, &input)
	if err != nil {
		return nil, errors.New("Product_Service.FindProductById : Failed to Decode Request Body : " + err.Error())
	}

	// Call Usecase
	res, err := s.repo.FindProductById(input)
	if err != nil {
		return nil, errors.New("Product_Service.FindProductById : Failed to Fetch Products : " + err.Error())
	}

	return res.(*proto.FindProductByIdResponse), nil
}