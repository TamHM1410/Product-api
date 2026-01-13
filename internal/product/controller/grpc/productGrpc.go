package grpc

import (
	"go-backend-api/internal/product/application/service"
	"go-backend-api/pb"
)

type ProductGrpcServer struct {
	pb.UnimplementedProductSimpleServiceServer
	service service.ProductService
}
