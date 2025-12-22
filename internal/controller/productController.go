package controller

// import (
// 	productservice "go-backend-api/internal/services/productService"
// 	"go-backend-api/response"

// 	"github.com/gin-gonic/gin"
// )

// type ProductController struct {
// 	productservice    *productservice.ProductService
// 	productRepository interface{}
// }

// func NewProductController() *ProductController {
// 	return &ProductController{}
// }

// // Thêm các method xử lý request liên quan đến Product ở đây
// func (pc *ProductController) GetProductByID(c *gin.Context) {
// 	// Dummy implementation
// 	id := c.Query("id")
// 	response.SuccessResponse(c, 1)
// 	if id == "1" {
// 		response.SuccessResponse(c, 1)
// 	}
// 	response.ErrorResponse(c, 404)
// }

// func (pc *ProductController) CreateProduct(c *gin.Context) {

// 	// payload:=c.Request.Body
// 	// pc.productservice.Create(payload)
// 	// Dummy implementation for creating a product

// }
