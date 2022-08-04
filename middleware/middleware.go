package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	tokenUsecase "ewallet/token/usecase"
)

func extractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenString := extractToken(ctx)

		token, err := tokenUsecase.Validate(tokenString)

		if err != nil {
			if strings.Contains("Token is expired", err.Error()) {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				ctx.Abort()
				return
			}
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			return
		} else {

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
