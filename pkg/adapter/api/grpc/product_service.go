package grpc

import (
	"context"
	"errors"

	usecase "github.com/product/pkg/usecase/product"
	proto "github.com/product/pkg/infrastructure/grpc/proto/product"
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