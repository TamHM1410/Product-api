package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {

	router := gin.Default()

	// v1 := router.Group("/api/v1")

	// handler := initialize.InitialProduct(db)

	// productHttp.ProductRouter(v1, handler)

	return router
}
