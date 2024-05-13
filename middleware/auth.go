package middleware

import (
	session "Game/sessions"
	"Game/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Auth(c *gin.Context) {
	if c.Request.URL.Path == "/login" {
		tools.Info(c.Request.URL.Path)
		c.Next()
	}

	store := session.NewSessionStore()
	sessions, err := store.Get(c.Request, "session")
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		c.String(404, "%s", err.Error())
		c.Abort()
		return
	}
	if sessions.IsNew {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Abort()
		return
	}
	username := c.Query("username")
	if val, ok := sessions.Values["user_phone"]; ok {
		if s, e := val.(string); e {
			if s == username {
				c.Next()
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
	c.AbortWithStatus(http.StatusUnauthorized)
}
