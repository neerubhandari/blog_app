package Config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// store the memory address of a "gorm.DB" object
var DB *gorm.DB

// struct for configuring
type Config struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// function returns a pointer to the "Config" object using the "&" operator
func BuildDBConfig() *Config {
	dbConfig := Config{
		Host:     "localhost",
		Port:     3306,
		User:     "sergey",
		Password: "sergey",
		DBName:   "photo_app",
	}
	return &dbConfig
}

// function takes a database configuration struct and returns a connection URL
func DbURL(dbConfig *Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
