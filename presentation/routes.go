package presentation

import (
	"go-webapp-practice/infrastructure"
	"go-webapp-practice/presentation/controller"
	"go-webapp-practice/presentation/middleware"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Routing Setup
func SetupRouter(r *gin.Engine) *gin.Engine {

	userController := controller.GetUserController()

	// CORS 設定
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// Static File Publishing
	r.Static("/assets", "./presentation/public/browser")
	// Publish index.html with root path
	r.StaticFile("/", "./presentation/public/browser/index.html")

	// APIグループ
	api := r.Group("/api")
	api.GET("/config", func(c *gin.Context) {
		auth0Config := infrastructure.LoadAUTH0Config()
		c.JSON(200, auth0Config)
	})

	// Endpoints requiring authentication
	authRequired := api.Group("/")
	authRequired.Use(middleware.Auth0Middleware(), middleware.FinalAuthMiddleware(userController))
	{
		// authRequired.POST("/users", userController.CreateUser)
		authRequired.GET("/users/:id", userController.GetUser)
	}

	// Returns index.html for requests that do not match any path
	r.NoRoute(func(c *gin.Context) {
		c.File("./public/browser/index.html")
	})

	return r
}
