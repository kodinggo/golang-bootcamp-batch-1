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
	db := NewMysql()

	rows, err := db.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatal(err.Error())
		}

		users = append(users, user)
	}

	jsonString, _ := json.MarshalIndent(users, "", "  ")
	log.Println(string(jsonString))
}
