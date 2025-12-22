package productRouter

import (
	"go-backend-api/internal/global"

	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine) {

	// public := router.Group(global.API_PREFIX)
	// {
	// 	public.GET("/product", controller.NewProductController().GetProductByID)
	// 	public.POST("/product", controller.NewProductController().CreateProduct)
	// }

	//
	private := router.Group(global.API_PREFIX)
	{
		// Add private routes here with authentication middleware
		private.GET("/post", func(c *gin.Context) {
			c.JSON(200, gin.H{ /// map string
				"message": "pong",
				"status":  200,
				"data":    nil,
			})
		})
	}
}
