package http

import (
	"fmt"
	"go-backend-api/internal/product/application/service"
	"go-backend-api/internal/product/application/service/dto"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) (res interface{}, err error, code int) {

	var req dto.ProductDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err, 400
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name is required"), 400
	}

	// Implementation for creating a product
	fmt.Println("req", req)

	createdProduct, err := h.service.CreateProduct(&req)

	if err != nil {
		return nil, err, 500
	}

	fmt.Println("createdProduct", createdProduct)

	return createdProduct, nil, 200
}

func (h *ProductHandler) GetProductByID(ctx *gin.Context) (res interface{}, err error, code int) {
	idParam := ctx.Param("id")
	fmt.Print(idParam, "id params")
	var id uint
	_, err = fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		return nil, fmt.Errorf("invalid id parameter"), 400
	}

	product, err := h.service.GetProductByID(id)
	if err != nil {
		return nil, err, 500
	}
	if product == nil {
		return nil, fmt.Errorf("product not found"), 404
	}

	return product, nil, 200
}
