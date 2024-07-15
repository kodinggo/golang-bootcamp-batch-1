package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	var (
		username = "root"
		password = "root"
		hostname = "localhost"
		port     = "3306"
		dbname   = "kodinggo"
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		username, password, hostname, port, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	// setup for manual direction and step
	var (
		direction = "up"
		step      = 2
	)

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	var n int
	if direction == "up" {
		n, err = migrate.ExecMax(db, "mysql", migrations, migrate.Up, step)
	} else {
		n, err = migrate.ExecMax(db, "mysql", migrations, migrate.Down, step)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Applied %d migrations!\n", n)
}
