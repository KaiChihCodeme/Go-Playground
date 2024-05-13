//CRUD
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbUser = "name"
	dbPassword = "111"
	dbName = "test"
)

func main () {
	// connect to DB
	// db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname)
	db, err := sql.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUser, dbPassword, dbName))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create table

	// CRUD tasks
}

func createTable() {

}

func inseert(){

}

func query(){

}

func update(){

}

func delete() {

}