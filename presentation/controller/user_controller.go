package controller

import (
	"net/http"
	"strconv"

	"go-webapp-practice/application"
	"go-webapp-practice/infrastructure/db"
	"go-webapp-practice/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *application.UserService
}

func GetUserController() *UserController {
	// TODO: use interface
	userRepository := repositories.NewUserRepository(db.DB)
	userService := application.NewUserService(userRepository)
	return NewUserController(userService)
}

func NewUserController(userService *application.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := uc.userService.CreateUser(req.Name, req.Email, req.Password)
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
