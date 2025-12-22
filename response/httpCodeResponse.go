package response

import "github.com/gin-gonic/gin"

// HTTPCodeResponse represents a standard HTTP response with a code and message.
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewResponseData creates a new HTTPCodeResponse with the given code and message.
func NewResponseData(code int, message string, data interface{}) *ResponseData {
	return &ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(HTTPStatusOK, NewResponseData(HTTPStatusOK, GetMessage(HTTPStatusOK), data))
}
func ErrorResponse(c *gin.Context, code int) {
	c.JSON(code, NewResponseData(code, GetMessage(code), nil))
}

type HandlerFunc func(c *gin.Context) (interface{}, error, int)

func WrapHandler(h HandlerFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		data, err, code := h(c)
		if err != nil {
			ErrorResponse(c, code)
			return
		}
		SuccessResponse(c, data)
	}
}
