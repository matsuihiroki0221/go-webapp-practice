package db

import (
	"go-webapp-practice/infrastructure"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg := infrastructure.LoadDBConfig()

	// データベースを作成
	CreateDatabase(cfg)

	// マイグレーションを実行
	RunMigration(cfg)

	// DB接続
	db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connection established successfully!")
}
