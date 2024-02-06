package database

import (
	"fmt"

	"github.com/Roninors/Expense_Tracker/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBname   string
}

var DB *gorm.DB

func ConnectDb(config *Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBname)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("successfully connected to database")
	// DB.Exec("ALTER TABLE users DROP CONSTRAINT idx_users_name")
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	fmt.Println("Successfuly migrated")

	return err
}
