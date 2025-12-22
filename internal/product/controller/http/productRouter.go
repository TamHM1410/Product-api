package http

import (
	"go-backend-api/response"

	"github.com/gin-gonic/gin"
)

func ProductRouter(rg *gin.RouterGroup, handler *ProductHandler) {
	productGroup := rg.Group("/products")

	productGroup.POST("/", response.WrapHandler(handler.CreateProduct))
	productGroup.GET("/:id", response.WrapHandler(handler.GetProductByID))
}
