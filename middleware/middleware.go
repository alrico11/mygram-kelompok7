package middleware

import (
	"net/http"
	"project2/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", "Token Not Found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Bearer tokentokentoken
		arrayToken := strings.Split(authHeader, " ")

		var tokenString string
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := NewService().ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", "Token Invalid!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", "Token invalid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["id_user"].(float64))

		if userID == 0 {
			response := helper.APIResponse("Unauthorized", "ID Not Found!")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", userID)
	}
}
