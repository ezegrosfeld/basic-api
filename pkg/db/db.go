package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Database   *sql.DB
	production bool = os.Getenv("ENVIRONMENT") == "production"
)

func InitializeDatabase() {
	url := os.Getenv("DB_URL")
	if production && url == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Panic(err)
	}

	// Ping the database to check connection
	err = db.Ping()
	if production && err != nil {
		log.Panic(err)
	}

	Database = db
}
