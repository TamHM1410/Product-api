package middleware

import (
	"go-backend-api/internal/services/userService"
	"go-backend-api/response"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		bearerToken := c.GetHeader("Authentication")

		if bearerToken == "" {
			response.ErrorResponse(c, 401)
			c.Abort()
			return
		}
		token := bearerToken[len("Bearer "):]

		authenticated := userService.AuthenticateUser(token)

		if !authenticated {
			response.ErrorResponse(c, 401)
			c.Abort()
			return
		}
		c.SetAccepted(token)
		c.Next()

	}
}
