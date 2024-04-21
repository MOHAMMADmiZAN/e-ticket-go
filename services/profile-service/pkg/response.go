package pkg // Package pkg APIResponse represents a standard API response structure.
import "github.com/gin-gonic/gin"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// RespondWithError handles error responses with a standard format.
func RespondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, APIResponse{Success: false, Error: err.Error()})
}

// RespondWithSuccess handles success responses with a standard format.
func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}, message string) {
	response := APIResponse{Success: true, Data: data}
	if message != "" {
		response.Message = message
	}
	c.JSON(statusCode, response)
}
