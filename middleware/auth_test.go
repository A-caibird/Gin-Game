package middleware

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func Test_Auth(t *testing.T) {
}

func ExampleAuth() {
	c := gin.New()
	c.Use(Auth)
	c.Run(":8080")
}
