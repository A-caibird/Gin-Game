package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors A middleware for Configuring CORS response headers and handling preflight request
func Cors(c *gin.Context) {
	// Configure CORS response headers
	c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	c.Header("Access-Control-Allow-Headers", "Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	// handle preflight request
	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Max-Age", "86400")
		c.Writer.WriteHeader(http.StatusOK)
		return
	}
	c.Next()
}
