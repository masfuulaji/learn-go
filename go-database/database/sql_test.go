package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `INSERT INTO "user" (id, name) VALUES ('U_001','Udin')`
	_, err := db.Exec(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name FROM "user"`
	rows, err := db.Query(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
	}
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := `SELECT id, name, email, balance, rating, birth_date, married, created_at FROM "user"`
	rows, err := db.Query(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, email string
		var name sql.NullString
		var balance int32
		var rating float32
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		if name.Valid {
			fmt.Println("Name :", name)
		}
		fmt.Println("Email :", email)
		fmt.Println("Balance :", balance)
		fmt.Println("rating :", rating)
		if birthDate.Valid {
			fmt.Println("Birth Date :", birthDate)
		}
		fmt.Println("Married :", married)
		fmt.Println("Created At :", createdAt)
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username FROM admin WHERE username = $1 AND password = $2 LIMIT 1"
	rows, err := db.Query(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "doni@gmail.com"
	comment := "Kerja Bagus"

	query := "INSERT INTO comments(email, comment) VALUES  ($1, $2) RETURNING id"
	var id int
	err := db.QueryRow(ctx, query, email, comment).Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data with id", id)
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin(ctx)
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES  ($1, $2) RETURNING id"
	var id int
	for i := 0; i < 10; i++ {
		email := "yono" + strconv.Itoa(i) + "@gmail.com"
		comment := "Kerja bagus " + strconv.Itoa(i)

		err = tx.QueryRow(ctx, query, email, comment).Scan(&id)
		if err != nil {
			fmt.Println(err)
			_ = tx.Rollback(ctx)
		}
		fmt.Println("Comment from", id)
	}

	err = tx.Commit(ctx)
	if err != nil {
		panic(err)
	}
}
