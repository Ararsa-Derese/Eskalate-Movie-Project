package middleware

import (
	"eskalate-movie-api/pkg/response"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					response.NewErrorResponse(
						"Internal server error",
						[]string{"internal server error"},
					),
				)
			}
		}()
		c.Next()
	}
}
