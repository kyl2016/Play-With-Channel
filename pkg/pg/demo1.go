package main

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

func main() {
	connStr := "host=localhost port=5432 user=lynxi password=lynxi dbname=ivs sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("users", "user_name", "role_id"))
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		Name   string
		RoleID int
	}

	users := []User{
		{"Kitty", 1},
		{"Bob", 0},
	}

	for _, user := range users {
		_, err = stmt.Exec(user.Name, int64(user.RoleID))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
