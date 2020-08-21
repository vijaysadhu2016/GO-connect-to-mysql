package main

import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main()  {
	fmt.Println("Go MySql Tutorial")
	db := dbConn()

	defer db.Close()
}