package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHOST = "DBHOST"
	DBUSER = "DBUSER"
	DBPASS = "DBPASS"
	DBNAME = "DBNAME"
)

var db *sql.DB

func main() {
	_, err := sql.Open("mysql", DBUSER+":"+DBPASS+"@tcp("+DBHOST+")/"+DBNAME+"")
	if err != nil {
		panic(err)
	}

	println("conectado!")
}
