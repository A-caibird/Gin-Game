package middleware

import (
	session "Game/sessions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"regexp"
)

func Auth(c *gin.Context) {
	// Determine whether verification can be skipped
	var excludeAuthPath []string
	excludeAuthPath = append(excludeAuthPath, "/login", "/signup", "/sms")
	for _, v := range excludeAuthPath {
		path := c.Request.URL.Path
		re := regexp.MustCompile(v)
		if re.MatchString(path) {
			c.Next()
			return
		}
	}
	//
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
	//
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
