package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

func TestHelloHandler(t *testing.T) {
	r := getRouter(true)

	r.GET("/hello", helloHandler())

	req, _ := http.NewRequest("GET", "/hello", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		return statusOK
	})
}
