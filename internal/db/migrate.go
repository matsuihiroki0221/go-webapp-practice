package db

import (
	"database/sql"
	"fmt"
	"go-webapp-practice/config"
	"log"
	"net"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// CreateDatabase はデータベースが存在しない場合に作成
func CreateDatabase(cfg *config.DBConfig) {

	db, err := sql.Open("mysql", cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
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

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", cfg.DBName))
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	log.Println("Database ensured:", cfg.DBName)
}

// RunMigrations はマイグレーションを実行
func RunMigrations(cfg *config.DBConfig) {

	db, err := sql.Open("mysql", cfg.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database for migration:", err)
	}
	defer db.Close()

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}

	m, err := migrate.NewWithDatabaseInstance(

		"file://internal/db/migrations", // マイグレーションファイルのパス
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
