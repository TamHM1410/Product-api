package grpc

import (
	"context"
	"go-backend-api/internal/product/application/service"
	"go-backend-api/pb"

	"google.golang.org/grpc"
)

type ProductGrpcServerHandler struct {
	pb.UnimplementedProductSimpleServiceServer
	service service.ProductService
}

func NewProductGrpcServerHandler(s *grpc.Server, service service.ProductService) *ProductGrpcServerHandler {
	return &ProductGrpcServerHandler{
		service: service,
	}
}

func (ps *ProductGrpcServerHandler) CreateOrUpdate(ctx context.Context, req *pb.ProductSimpleDTO) (*pb.ProductSimpleResponse, error) {

	if req == nil {
		return nil, nil
	}

	return &pb.ProductSimpleResponse{
		Id:          1,
		Name:        req.Name,
		Description: req.Description,
		Height:      req.Height,
		Width:       req.Width,
		Depth:       req.Depth,
	}, nil
}
