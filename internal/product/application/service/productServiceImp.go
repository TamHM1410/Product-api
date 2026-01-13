package service

import (
	rabbitmq "go-backend-api/config/rabbitMQ"
	"go-backend-api/internal/product/application/service/dto"
	product "go-backend-api/internal/product/domain/model/entity"
	"go-backend-api/internal/product/domain/repository"
	"go-backend-api/pb"

	"google.golang.org/protobuf/encoding/protojson"

	amqp "github.com/rabbitmq/amqp091-go"
)

type productService struct {
	productRepo *repository.ProductRepo
	rabbit      *amqp.Connection
}

func (ps *productService) CreateProduct(payload *dto.ProductDTO) (*product.ProductModel, error) {
	var toSave product.ProductModel

	toSave.Name = payload.Name
	toSave.Description = payload.Description
	toSave.Price = payload.Price

	createdProduct, err := ps.productRepo.Create(&toSave)

	return createdProduct, err
}
func (ps *productService) GetProductByID(id uint) (*product.ProductModel, error) {
	// Implementation for getting a product by ID
	product, err := ps.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
func (ps *productService) CreateOrUpdate(payload *pb.ProductSimpleDTO) (*product.ProductModel, error) {
	// Implementation for creating or updating a product
	var toSave product.ProductModel
	var typeOfFunc = "Create"

	if payload.Id != 0 {
		typeOfFunc = "Update"
	}

	toSave.ID = uint(payload.Id)
	toSave.Name = payload.Name
	toSave.Description = payload.Description
	toSave.Height = payload.Height
	toSave.Width = payload.Width
	toSave.Depth = payload.Depth

	savedProduct, err := ps.productRepo.CreateOrUpdate(&toSave)

	if err != nil {
		return nil, err
	}

	if typeOfFunc == "Create" {
		
		payload.Id =int32(savedProduct.ID)

		jsonBytes, err := protojson.Marshal(payload)
		
		if err != nil {
			return nil, err
		}

		rabbitmq.SendMessage(ps.rabbit, "inventory", string(jsonBytes))

	}

	return savedProduct, nil
}

func NewProductService(repo *repository.ProductRepo, rabbit *amqp.Connection) ProductService {
	return &productService{
		productRepo: repo,
		rabbit:      rabbit,
	}
}
