package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

//store limiter per IP
var clients = make(map[string]*rate.Limiter)
var mu sync.Mutex

//get limiter for each IP
func getLimiter(ip string) *rate.Limiter{
    mu.Lock()
    defer mu.Unlock()
    
    limiter, exists := clients[ip]
    if !exists{
        //5 req per sec
        limiter = rate.NewLimiter(5, 10)
        clients[ip] = limiter
    }
    return limiter
}


//middleware
func RateLimitMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        
        ip := c.ClientIP()
        limiter := getLimiter(ip)
        
        if !limiter.Allow(){
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error":"too many requests",
            })
            c.Abort()
            return 
        }
        c.Next()
    }
}