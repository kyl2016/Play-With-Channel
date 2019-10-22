package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	connStr := "host=192.168.11.187 port=5432 user=lynxi password=lynxi dbname=ivs sslmode=disable"
	//connStr := "host=192.168.11.187 port=5432 user=lynxi password=lynxi dbname=ivs sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Eigen string
	}

	var dest []Data
	err = db.Raw("select eigen from repo_faces where repo_id = 17 and uuid = '5d63c58ecad04446093de608'").Scan(&dest).Error
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("/tmp/faceEigen", []byte(dest[0].Eigen), os.ModePerm)

	fmt.Println(dest)
}
