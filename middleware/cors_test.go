package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_cors(t *testing.T) {
	c := gin.Default()
	c.GET("/hello", Cors)
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	assert.Equal(t, w.Header().Get("Access-Control-Allow-Credentials"), "true")
	fmt.Printf("%#v", w.Result())
}
