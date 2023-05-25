package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func GetConnection() *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), "postgres://postgres:root@localhost:5432/golang")

	if err != nil {
		fmt.Println("Unable to connect to database", err)
		os.Exit(1)
	}
	//defer dbpool.Close()
	return dbpool
}
