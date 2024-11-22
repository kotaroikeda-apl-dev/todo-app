package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

var DB *gorm.DB // グローバル変数を追加

func ConnectDB() error {
    host := os.Getenv("DATABASE_HOST")
    user := os.Getenv("DATABASE_USER")
    password := os.Getenv("DATABASE_PASSWORD")
    dbName := os.Getenv("DATABASE_NAME")
    port := "5432"

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port)
    
    // デバッグ用に DSN を出力
    fmt.Println("Connecting to database with DSN:", dsn)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return fmt.Errorf("failed to connect to database: %w", err)
    }

    DB = db // グローバル変数に代入
    return nil
}
