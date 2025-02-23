package main

import (
	"go-webapp-practice/internal/handlers"
	"go-webapp-practice/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter ルーティングのセットアップ
func SetupRouter(r *gin.Engine) *gin.Engine {

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 静的ファイルの公開
	r.Static("/assets", "./public/assets")
	// ルートパスでindex.htmlを公開
	r.StaticFile("/", "./public/index.html")

	// APIグループ
	api := r.Group("/api")

	// 認証が必要なエンドポイント
	authRequired := api.Group("/")
	authRequired.Use(middleware.Auth0Middleware())
	{
		authRequired.GET("/users", handlers.GetUsersHandler)
	}

	// 該当しないリクエストパスにはindex.htmlを返却
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return r
}
