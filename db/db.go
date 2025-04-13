package db

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func Init() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "secret",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "mydb",
		AllowNativePasswords: true,
	}

	var err error

	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
