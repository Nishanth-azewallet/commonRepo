package common

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// ConnectDB initializes the DB connection
func ConnectDB(dbName string) error {
	dsn := "root:password@tcp(127.0.0.1:3306)/" + dbName // Change this as needed
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Verify the connection to ensure it's working
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	log.Println("Database connected successfully!")
	return nil
}

// CloseDB closes the DB connection
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
