// pkg/middleware/swagger_host.go
package middleware

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// SwaggerHost sets the Swagger host dynamically based on the environment variable.
func SwaggerHost() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := os.Getenv("SWAGGER_HOST")
		if host == "" {
			host = "localhost:8081"
			log.Printf("SWAGGER_HOST environment variable not set. Defaulting to %s", host)
		}
		c.Set("swagger_host", host)
		c.Next()
	}
}
