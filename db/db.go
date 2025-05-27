package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}
	Db, _ = sql.Open("mysql", cfg.FormatDSN())

	if err := Db.Ping(); err != nil {
		log.Fatal("Ping error:", err)
	}
}
