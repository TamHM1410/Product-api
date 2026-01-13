package initialize

import (
	"go-backend-api/internal/product/application/service"
	productgrpc "go-backend-api/internal/product/controller/grpc"
	"go-backend-api/internal/product/domain/repository"

	amqp "github.com/rabbitmq/amqp091-go"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

//	func InitialProduct(db *gorm.DB) *product.ProductHandler {
//		repo := repository.NewProductRepo(db)
//		service := service.NewProductService(repo)
//		handler := product.NewProductHandler(service)
//		return handler
//	}
func InitialProductGRPC(s *grpc.Server, db *gorm.DB, rabbit *amqp.Connection) *productgrpc.ProductGrpcServerHandler {
	repo := repository.NewProductRepo(db)
	service := service.NewProductService(repo, rabbit)
	return productgrpc.NewProductGrpcServerHandler(s, service)
}
