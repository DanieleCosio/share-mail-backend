package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	User     string
	Password string
	Database string
	Host     string
	Port     string
}

var ConfigCreated bool = false
var config Config

func CreateConfig(user string, password string, database string, host string, port string) {
	ConfigCreated = true
	config = Config{
		User:     user,
		Password: password,
		Database: database,
		Host:     host,
		Port:     port,
	}
}

func GetSqlConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.User,
		config.Password,
		config.Database,
		config.Host,
		config.Port,
	)

	return sql.Open("postgres", connStr)
}
