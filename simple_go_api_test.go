package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTest() *gin.Engine {
	gin.SetMode(gin.TestMode)
	req := gin.Default()
	return req
}

func performRequestTest(req http.Handler, method, path string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	req.ServeHTTP(w, r)
	return w
}

func TestPingHandler(t *testing.T) {
	// initialize router
	router := setupTest()
	router.GET("/ping", PingHandler)

	// perform request
	resp := performRequestTest(router, "GET", "/ping")

	// assert response
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.JSONEq(t, `{"message":"pong"}`, resp.Body.String())
}
