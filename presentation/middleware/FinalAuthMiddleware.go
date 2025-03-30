package middleware

import (
	"go-webapp-practice/domain/models"
	"go-webapp-practice/presentation/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FinalAuthMiddleware(userController *controller.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		providerId, exists := c.Get("userId")
		if !exists || providerId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No valid authentication provider"})
			c.Abort()
			return
		}
		providerName, exists := c.Get("providerName")
		if !exists || providerId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No valid authentication provider"})
			c.Abort()
			return
		}

		// `providerID` でユーザーを検索
		user, err := userController.UserService.GetUserByProviderID(providerId.(string), providerName.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
			c.Abort()
			return
		}

		// ユーザーが存在しない場合、新規登録
		if user == nil {
			newUser := &models.User{
				ProviderId:   providerId.(string),
				ProviderName: providerName.(string),
			}

			err = userController.UserService.CreateUser(newUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
