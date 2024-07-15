package main

import (
	"database/sql"
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

	user := User{
		Name:      "John Doe",
		Email:     "john@doe.com",
		CreatedAt: time.Now(),
	}

	result, err := db.Exec(
		"INSERT INTO users (name, email, created_at) VALUES (?, ?, ?)",
		user.Name, user.Email, user.CreatedAt,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	id, _ := result.LastInsertId()
	log.Println("Last ID", id)
}
