package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

func main() {
	url := "host=localhost port=5432 user=lynxi password=lynxi dbname=test sslmode=disable"
	db, err := gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	db.SingularTable(false)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// create table
	db.AutoMigrate(&User{})

	// batch insert multiple data
	users := []interface{}{
		&User{1, "Kitty", time.Now().AddDate(-30, 0, 0)},
		&User{2, "Bob", time.Now().AddDate(-30, 0, 0)},
		&User{3, "Lili", time.Now().AddDate(-30, 0, 0)},
	}
	BatchCreate(db, users)

	fmt.Println("completed")
}

type User struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Birthday time.Time
}
