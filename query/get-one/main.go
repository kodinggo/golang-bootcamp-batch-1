package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func NewMysql() *sql.DB {
	var (
		username = "root"
		password = "root"
		hostname = "localhost"
		port     = "3306"
		dbname   = "kodinggo"
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, hostname, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// initiate db connection
	db := NewMysql()

	// initiate struct for query
	query := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "Alice",
	}

	// perform query to database
	row := db.QueryRow(
		`SELECT id, name, email, created_at FROM users WHERE id = ? AND name = ?`,
		query.ID,
		query.Name,
	)

	// initiate struct for result
	var user User

	// scan the result to -> struct
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Fatal(err.Error())
	}

	// print the result
	jsonString, _ := json.MarshalIndent(user, "", "  ")
	log.Println(string(jsonString))
}
