package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 数据库实例
var DB *gorm.DB

// Initializing the database.
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values.")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPasswd, dbHost, dbPort, dbName)
	// GORM 定义了这些日志级别：Silent、Error、Warn、Info
	DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
		// DisableForeignKeyConstraintWhenMigrating: true,
		// SkipDefaultTransaction:                   true,
	})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbDriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected the %s database", dbDriver)
	}

	// database connection pool settings.
	// refer to https://www.alexedwards.net/blog/configuring-sqldb
	sqlDB, _ := DB.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(64)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(64)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// database migration
	DB.Debug().AutoMigrate(&Shop{}, &Order{}, &Delivery{}, &Detail{}, &Payment{})
}
