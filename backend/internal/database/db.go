package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host string
	Port string
	User string 
	Password string
	DBName string
}

func NewConnection(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
							config.Host, config.Port, config.User, config.Password, config.DBName)
							
	return sql.Open("postgres", connStr)
}