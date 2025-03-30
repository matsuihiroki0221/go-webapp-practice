package db

import (
	"database/sql"
	"fmt"
	"go-webapp-practice/infrastructure"
	"log"
	"net"

	"github.com/golang-migrate/migrate/v4"
	migratemysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	cfg := infrastructure.LoadDBConfig()

	RunMigration(cfg)

	db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connection established successfully!")
}

// RunMigration はマイグレーションを実行
func RunMigration(cfg *infrastructure.DBConfig) {

	db, err := sql.Open("mysql", cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database for migration:", err)
	}
	defer db.Close()

	// 接続テスト
	if err := db.Ping(); err != nil {
		if opErr, ok := err.(*net.OpError); ok {
			log.Fatalf("Network error: %v", opErr)
		} else {
			log.Fatalf("Failed to ping database: %v", err)
		}
	}

	// DB が存在しない場合は作成
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DBName))
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	driver, err := migratemysql.WithInstance(db, &migratemysql.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(

		"file://infrastructure/db/migration", // マイグレーションファイルのパス
		"mysql", driver,
	)
	if err != nil {
		log.Fatal("Failed to initialize migration:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database migration completed successfully!")
}
