package response

const (
	HTTPStatusOK                  = 200
	HTTPStatusBadRequest          = 400
	HTTPStatusUnauthorized        = 401
	HTTPStatusForbidden           = 403
	HTTPStatusNotFound            = 404
	HTTPStatusInternalServerError = 500
)

// /message
var msg = map[int]string{
	HTTPStatusOK:                  "OK",
	HTTPStatusBadRequest:          "Bad Request",
	HTTPStatusUnauthorized:        "Unauthorized",
	HTTPStatusForbidden:           "Forbidden",
	HTTPStatusNotFound:            "Not Found",
	HTTPStatusInternalServerError: "Internal Server Error",
}

func GetMessage(code int) string {
	if message, exists := msg[code]; exists {
		return message
	}
	return "Unknown Status Code"
}
