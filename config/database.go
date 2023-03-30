package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/yosikez/item-api/model"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

func (c *DatabaseConfig) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=%s TimeZone=%s", c.Host, c.Username, c.Password, c.Name, c.Port, c.SSLMode, c.TimeZone)
}

func InitDB() (*gorm.DB, error) {
	dbConfig := loadDatabase()
	dsn := dbConfig.dsn()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Item{})

	return db, nil
}

func loadDatabase() *DatabaseConfig {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error load .env file")
	}

	dbConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		TimeZone: os.Getenv("DB_TIMEZONE"),
	}

	return dbConfig
}
