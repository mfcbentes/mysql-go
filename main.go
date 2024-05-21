package main

import (
	"database/sql"
	"os"

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
	_, err := sql.Open("mysql", os.Getenv(DBUSER)+":"+os.Getenv(DBPASS)+"@tcp("+os.Getenv(DBHOST)+")/"+os.Getenv(DBNAME)+"")
	if err != nil {
		panic(err)
	}

	println("conectado!")
}
