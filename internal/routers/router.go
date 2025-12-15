package routers

import (
	"fmt"
	"go-backend-api/internal/controller"
	"go-backend-api/response"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware before request")
		// Do something before request
		response.ErrorResponse(c, 403)

		c.Abort()
		c.Next()

		return // Stop processing the request further

		// Do something after request
		fmt.Println("Middleware after request")
	}
}

func NewRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{ /// map string
				"message": "pong",
				"status":  200,
				"data":    nil,
			})
		})
		///user api
		v1.GET("/user", AA(), controller.NewUserController().GetUser)
	}

	return r
}
