package controller

import (
	application_service "go-webapp-practice/application_service"
	"go-webapp-practice/infrastructure/db"
	"go-webapp-practice/infrastructure/repositories"
)

var (
	UserCtrller *UserController
	// 他のコントローラーもここに追加
)

func Initializecontroller() {

	// 依存関係の注入
	userRepository := repositories.NewUserRepository(db.DB)
	userService := application_service.NewUserService(userRepository)
	UserCtrller = NewUserController(userService)
	// 他のコントローラーもここで初期化
}
