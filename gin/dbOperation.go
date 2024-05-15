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

func insert(){
	db, err := conntection()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	_, err := db.Exec("Insert into table1 (id, name) values (?, ?)", id, name)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func query(){
	db, err := GetConnection()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	rows, err := db.Query("Select * from table where id = ?")
	if err != nil {
		fmt.Println(err.Error())	
	}

	defer rows.Close()

	var dbResult models.GetDbResult
	for rows.Next() {
		if err := rows.scan(&dbResult.id, &dbResult.Name); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err.Error())
	}

	if len(dbResult) == 0 {
		return nil
	}

}

func update(){
	db, err := connection()
	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	res, err := db.Exec("Update table1 set id=?, name=?", id, name)
	if err != nil {
		fmt.Println(err.Error())
	}

	isAffected, err := res.RowAffected()
	if err != nil {
		fmt.Println(err.Error())
	}

	if isAffected == 0 {
		return fmt.Errorf("Failed to update the db")
	}

}

func delete() {
	db.Exec("DELETE FROM table1 where id = ?", id)
}