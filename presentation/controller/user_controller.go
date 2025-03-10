package controller

import (
	"net/http"
	"strconv"

	"go-webapp-practice/application"
	"go-webapp-practice/domain/models"
	"go-webapp-practice/infrastructure/db"
	"go-webapp-practice/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *application.UserService
}

func GetUserController() *UserController {
	userRepository := repositories.NewUserRepository(db.DB)
	userService := application.NewUserService(userRepository)
	return NewUserController(userService)
}

func NewUserController(userService *application.UserService) *UserController {
	return &UserController{userService: userService}
}

// DTO for creating a user
type CreateUserRequest struct {
	Auth0Id string `json:"auth0_id" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

// DTOからドメインモデルへ変換するメソッド
func (req *CreateUserRequest) ToDomain() *models.User {
	return models.NewUser(req.Auth0Id, req.Name)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := req.ToDomain()

	err := uc.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uc *UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := uc.userService.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
