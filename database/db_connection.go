package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Database struct {
	Config     mysql.Config
	Connection *sql.DB
}

var db *Database

type TxFunc func(context.Context, *sql.Tx) error

func InitializeDbConnection() {
	fmt.Println("Initialize database connection...")

	cfg := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "echopost",
	}

	var err error
	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Successfully connected to database!")

	db = &Database{
		Config:     cfg,
		Connection: conn,
	}

	defer conn.Close()
}

func NewContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ctx
}

func ExecuteInTransaction(ctx context.Context, fn TxFunc) error {
	tx, txErr := db.Connection.BeginTx(ctx, nil)
	if txErr != nil {
		return txErr
	}

	defer tx.Rollback()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}
