package basic

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" // memanggil method init
)

func TestConnectionDB(t *testing.T) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/filament_app")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(5)                   // connection min 5
	db.SetMaxOpenConns(10)                  // connection max 10
	db.SetConnMaxIdleTime(5 * time.Minute)  // tiap 5 menit idle maka close connection
	db.SetConnMaxLifetime(60 * time.Minute) // tiap 60 menit maka connection akan diperbarui
}

func TestSQLCommand(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// digunakan untuk yg tidak membutuhkan hasil
	result, err := db.ExecContext(
		ctx,
		"INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		"testB",
		"testB@email.com",
		"password",
		time.Now(),
		time.Now(),
	)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("SUCCESS INSERT DATA", id)
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// untuk menggembalikan result
	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id uint64
		var name, email, password string
		var rememberToken sql.NullString
		var emailVerifiedAt sql.NullTime
		var createdAt, updatedAt time.Time

		err := rows.Scan(
			&id,
			&name,
			&email,
			&emailVerifiedAt,
			&password,
			&rememberToken,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			panic(err)
		}
		fmt.Println("id:", id, "name:", name, "email:", email)
	}

	fmt.Println("SUCCESS SELECT DATA")
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	// digunakan untuk query yang sama dengan param yg beda
	statement, err := db.PrepareContext(
		ctx,
		"INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
	)
	if err != nil {
		panic(err)
	}

	defer statement.Close()
	for i := 0; i < 10; i++ {
		result, err := statement.ExecContext(ctx,
			fmt.Sprintf("test%d", i),
			fmt.Sprintf("test%d", i)+"@email.com",
			"password",
			time.Now(),
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
		lastId, _ := result.LastInsertId()
		fmt.Println("Check LastId", lastId)
	}
}

func TestTransactionDB(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	for i := 0; i < 10; i++ {
		result, err := tx.ExecContext(ctx,
			"INSERT INTO users(name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
			fmt.Sprintf("test%d", i),
			fmt.Sprintf("test%d", i)+"@email.com",
			"password",
			time.Now(),
			time.Now(),
		)
		if err != nil {
			panic(err)
		}
		id, _ := result.LastInsertId()
		fmt.Println("last id", id)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}
}
