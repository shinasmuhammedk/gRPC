package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc{
    return func(c *gin.Context) {
        
        authHeader := c.GetHeader("Authorization")
        if authHeader == ""{
            c.JSON(http.StatusUnauthorized, gin.H{
                "error":"missing token",
            })
            c.Abort()
            return 
        }
        
        tokenStr := strings.Split(authHeader, " ")
        if len(tokenStr) != 2{
            c.JSON(http.StatusUnauthorized, gin.H{
                "error":"invalid token format",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}