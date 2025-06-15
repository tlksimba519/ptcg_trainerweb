package middlewares

import (
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tlksimba519/ptcg_trainerweb/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Missing token")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			utils.Error(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}
		tokenStr := parts[1]

		userID, _, err := utils.ValidateToken(tokenStr)
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
