package middleware

import (
	"eskalate-movie-api/pkg/response"
	"eskalate-movie-api/pkg/security"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.NewErrorResponse(
					"Missing or invalid Authorization header",
					[]string{"unauthorized"},
				),
			)
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := security.ParseJWT(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.NewErrorResponse(
					"Invalid token",
					[]string{"unauthorized"},
				),
			)
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
