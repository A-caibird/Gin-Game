package session

import (
	"Game/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSession(t *testing.T) {
	store1 := NewSessionStore()
	c := gin.Default()
	// Simulate browser requests
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()
	//
	c.GET("/hello", middleware.Cors, func(c *gin.Context) {
		session, _ := store1.Get(c.Request, "session")
		//var wg sync.WaitGroup
		//wg.Add(1)
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()
		//	session.Values["name"] = "fdasdfdsa"
		//}(&wg)
		//wg.Add(1)
		//go func(wg *sync.WaitGroup) {
		//	defer wg.Done()
		//	session.Values["name"] = "fasdfas23541235a"
		//	print("fdasdfdas")
		//}(&wg)
		//wg.Wait()
		if val, ok := session.Values["name"]; ok {
			if v, o := val.(string); o {
				println(v)
			}
		} else {
			println("request with carry cookie ,authentication fail")
		}
		session.Values["name"] = "Lian"
		session.Save(c.Request, c.Writer)
		c.Next()
	})
	c.ServeHTTP(w, req)
	//cookie := w.Header().Get("Set-Cookie")
	////fmt.Printf("%#v", cookie)
	//req = httptest.NewRequest(http.MethodGet, "/hello", nil)
	//req.Header.Set("Cookie", cookie)
	//w = httptest.NewRecorder()
	//c.ServeHTTP(w, req)
	//:fasdfasdfas
	//[GIN] 2024/05/07 - 21:07:50 | 200 |      173.75µs |       192.0.2.1 | GET      "/hello"
	//Lian
	//[GIN] 2024/05/07 - 21:07:50 | 200 |      72.666µs |       192.0.2.1 | GET      "/hello"
	//--- PASS: TestSession (0.00s)
}
