package basic

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/filament_app?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)                   // connection min 5
	db.SetMaxOpenConns(10)                  // connection max 10
	db.SetConnMaxIdleTime(5 * time.Minute)  // tiap 5 menit idle maka close connection
	db.SetConnMaxLifetime(60 * time.Minute) // tiap 60 menit maka connection akan diperbarui

	return db
}
