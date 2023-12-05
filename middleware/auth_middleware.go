package middleware

import (
	"File-Management-System/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		refreshToken := c.Request.Header.Get("refresh-token")
		if clientToken == "" || refreshToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			refreshClaims, err := helpers.ValidateToken(refreshToken)
			if err != "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err})
				c.Abort()
				return
			} else {
				helpers.UpdateAllTokens(clientToken, refreshToken, refreshClaims.Uid)
				c.Set("email", refreshClaims.Email)
				c.Set("name", refreshClaims.Name)
				c.Set("uid", refreshClaims.Uid)
				c.Next()
			}

		} else {
			c.Set("email", claims.Email)
			c.Set("name", claims.Name)
			c.Set("uid", claims.Uid)
			c.Next()
		}

	}
}
