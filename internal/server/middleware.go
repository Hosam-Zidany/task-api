package server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if h == "" || !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization required"})
			return
		}
		tokenString := strings.TrimPrefix(h, "Bearer ")
		tok, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{},
			func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(AppConfig.JWTSecret), nil
			})
		if err != nil || !tok.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		claims, ok := tok.Claims.(*jwt.RegisteredClaims)
		if !ok || claims.Subject=="" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			return
		}
		uid,err := strconv.ParseUint(claims.Subject,10,64)
		if err!=nil{
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid subject in token"})
			return	
		}
		c.Set("userID",uint(uid))
		c.Next()
	}
}
