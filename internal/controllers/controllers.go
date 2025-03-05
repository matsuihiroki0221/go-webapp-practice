package controllers

import (
	"go-webapp-practice/internal/db"
	"go-webapp-practice/internal/repositories"
	"go-webapp-practice/internal/services"
)

var (
	UserCtrller *UserController
	// 他のコントローラーもここに追加
)

func InitializeControllers() {

	// 依存関係の注入
	userRepository := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepository)
	UserCtrller = NewUserController(userService)
	// 他のコントローラーもここで初期化
}
