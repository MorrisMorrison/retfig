package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/MorrisMorrison/retfig/config"
	"github.com/go-sql-driver/mysql"
)

type Connection struct {
	Config   mysql.Config
	Database *sql.DB
}

type TxFunc func(context.Context, *sql.Tx) error

func NewConnection() *Connection {
	fmt.Println("Initialize database connection...")

	cfg := mysql.Config{
		User:      config.GetEnv("RETFIG_MYSQL_USER", "user"),
		Passwd:    config.GetEnv("RETFIG_MYSQL_PASSWORD", "password"),
		Net:       "tcp",
		Addr:      config.GetEnv("RETFIG_MYSQL_HOST", "127.0.0.1:3306"),
		DBName:    config.GetEnv("RETFIG_MYSQL_DATABASE_NAME", "retfig"),
		ParseTime: true,
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

	return &Connection{
		Config:   cfg,
		Database: conn,
	}
}

func (connection *Connection) CreateContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return ctx
}

func (connection *Connection) ExecuteInTransaction(ctx context.Context, fn TxFunc) error {
	tx, txErr := connection.Database.BeginTx(ctx, nil)
	if txErr != nil {
		return txErr
	}

	defer tx.Rollback()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	return tx.Commit()
}
