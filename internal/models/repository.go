package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"nononsensecode.com/book/internal/config"
)

var db *sql.DB

func InitDB(config *config.Config) {
	var err error
	switch(config.DatabaseConfig.EngineName) {
	case "sqlite3":
		db, err = sql.Open("sqlite3", config.DatabaseConfig.Filename)
		if err != nil {
			panic(err)
		}
		initSqlite()
	case "postgres":
		datasource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s%s", 
		config.DatabaseConfig.User, config.DatabaseConfig.Password,
		config.DatabaseConfig.Host, config.DatabaseConfig.Port,
		config.DatabaseConfig.DatabaseName, config.DatabaseConfig.DBOptions)
		db, err = sql.Open("postgres", datasource)
		if err != nil {
			panic(err)
		}
	default:
		log.Printf("There is no database type %s\n", config.DatabaseConfig.EngineName)
		os.Exit(2)
	}
}

func NewBookRepository(config *config.Config) (BookRepository, error) {
	var repo BookRepository
	switch(config.DatabaseConfig.EngineName) {
	case "sqlite3":
		repo = SqliteBookRepository{}
	case "postgres":
		repo = PostgresqlBookRepository{}
	default:
		return nil, fmt.Errorf("there is no database named %s", config.DatabaseConfig.EngineName)
	}
	return repo, nil
}