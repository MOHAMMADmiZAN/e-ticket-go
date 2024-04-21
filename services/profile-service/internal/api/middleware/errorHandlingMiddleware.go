package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"runtime/debug"
)

var logger *zap.Logger

func init() {
	// Initialize a production logger with JSON formatting.
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic("Cannot initialize Zap logging: " + err.Error())
	}
}

// ErrorHandlingMiddleware intercepts and manages errors and panics globally.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic and stack trace at the ERROR level.
				logger.Error("Unhandled error or panic caught by middleware",
					zap.Any("error", r),
					zap.String("stack", string(debug.Stack())),
				)

				// Provide a generic error response to maintain security and prevent leakage of sensitive information.
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "An unexpected error occurred. Please try again later.",
				})

				// Abort processing to ensure no further handlers are executed.
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		// Proceed with the next middleware or handler.
		c.Next()
	}
}
