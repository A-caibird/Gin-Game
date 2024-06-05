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
		// bad request
		c.String(400, "%s", err.Error())
		c.Abort()
		return
	}
	//
	if sessions.IsNew {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//
	//if val, ok := sessions.Values["ID"]; ok {
	//	if s, e := val.(string); e {
	//		if s == "1" {
	//			c.Next()
	//		} else {
	//			c.AbortWithStatus(http.StatusUnauthorized)
	//		}
	//	}
	//}
}
