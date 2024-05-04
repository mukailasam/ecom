package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gopkg.in/boj/redistore.v1"
)

func DatabaseConnection(host, port, user, dbname, password, sslmode string) (*sql.DB, error) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	return db, nil
}

func RediStore(size int, network string, host string, port string, password string, secretKey []byte) (*redistore.RediStore, error) {
	rdStore, err := redistore.NewRediStore(size, network, fmt.Sprintf("%s:%s", host, port), password, secretKey)
	if err != nil {
		return nil, err
	}

	return rdStore, nil
}
