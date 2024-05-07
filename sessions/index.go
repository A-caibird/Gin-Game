package sessions

import (
	"github.com/gorilla/sessions"
	"sync"
)

var store *sessions.CookieStore
var once sync.Once

// NewSessionStore return session store ,default using cookie storage engine
func NewSessionStore() *sessions.CookieStore {
	once.Do(func() {
		store = sessions.NewCookieStore([]byte("acaibird.com"))
	})
	return store
}
