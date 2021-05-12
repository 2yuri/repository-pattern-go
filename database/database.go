package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/hyperyuri/repository-pattern-go/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	DbHost := os.Getenv("DB_HOST")
	DbPort := os.Getenv("DB_PORT")
	DbUser := os.Getenv("DB_USER")
	DbName := os.Getenv("DB_NAME")
	DbSSlMode := os.Getenv("DB_SSL_MODE")
	DbPass := os.Getenv("DB_PASS")
	DbMaxIddleConns, _ := strconv.Atoi(os.Getenv("DB_MAX_IDDLE_CONNS"))
	DbMaxOpensConns, _ := strconv.Atoi(os.Getenv("DB_MAX_OPENS_CONNS"))

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", DbHost, DbPort, DbUser, DbName, DbSSlMode, DbPass)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
		log.Fatal("Error: ", err)
	}

	db = database

	config, _ := database.DB()

	config.SetMaxIdleConns(DbMaxIddleConns)
	config.SetMaxOpenConns(DbMaxOpensConns)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunAutoMigrations(db)
}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
