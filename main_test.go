package main

import (
	"go-webapp-practice/infrastructure/db"
	"go-webapp-practice/presentation"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test for GET /healthcheck
func TestHealthCheckRoute(t *testing.T) {
	db.Init()

	router := gin.Default()
	presentation.SetupRouter(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}
