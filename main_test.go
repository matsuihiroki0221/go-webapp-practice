package main

import (
	"go-webapp-practice/internal"
	"go-webapp-practice/internal/controllers"
	"go-webapp-practice/internal/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for GET /healthcheck
func TestHealthCheckRoute(t *testing.T) {
	db.Init()

	// 依存関係の注入
	controllers.InitializeControllers()

	router := gin.Default()
	internal.SetupRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}
