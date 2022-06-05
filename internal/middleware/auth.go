package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/worldwidepaniel/ria-course-crud/internal/utils"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Token"]
		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "no token",
			})
			return
		}
		if !utils.IsJWTValid(token[0]) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}
		c.Next()
	}
}
