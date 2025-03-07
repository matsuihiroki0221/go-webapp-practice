package presentation

import (
	"go-webapp-practice/presentation/controller"
	"go-webapp-practice/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// Routing Setup
func SetupRouter(r *gin.Engine) *gin.Engine {

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Static File Publishing
	r.Static("/assets", "./public/assets")
	// Publish index.html with root path
	r.StaticFile("/", "./public/index.html")

	// APIグループ
	api := r.Group("/api")

	// Endpoints requiring authentication
	authRequired := api.Group("/")
	authRequired.Use(middleware.Auth0Middleware())
	{
		authRequired.POST("/users", controller.UserCtrller.CreateUser)
		authRequired.GET("/users/:id", controller.UserCtrller.GetUser)
	}

	// Returns index.html for requests that do not match any path
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/index.html")
	})

	return r
}
