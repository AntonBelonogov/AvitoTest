package middleware

import (
	"AvitoTest/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header"})
			ctx.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		ctx.Request.Header.Add("user_id", strconv.Itoa(int(claims.UserID)))
		ctx.Next()
	}
}
