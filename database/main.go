package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// db env
	var (
		username = "root"
		password = "password"
		database = "mydb"
		host     = "localhost"
	)

	// Build connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8mb4",
		username,
		password,
		host,
		database,
	)

	// Connect to server db
	dbConn, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect mysql, error: %v", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("failed to ping mysql, error: %v", err)
	}

	fmt.Println("Connected!")

	var v string
	fmt.Scan(&v)
}
