package middleware

import (
	session "Game/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Auth(c *gin.Context) {
	if c.Request.URL.Path == "/login" {
		c.Next()
	}
	store := session.NewSessionStore()
	sessions, err := store.Get(c.Request, "session")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		c.String(404, "%s", err.Error())
	}
	if sessions.IsNew {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	username := c.Query("username")
	if val, ok := sessions.Values["username"]; ok {
		if s, e := val.(string); e {
			if s == username {
				c.Next()
			}
		}
	}
	c.Writer.WriteHeader(http.StatusUnauthorized)
	c.Abort()
}
