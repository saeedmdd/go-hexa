package utils

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

// GetPostgresDsn generates
func GetPostgresDsn(host, user, password, port, database, sslmode string) string {

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, database, sslmode)
}

// PostgresConnection connects to postgres and if it does't connect panics
func PostgresConnection(host, user, password, port, database, sslmode string) (*sql.DB, error) {
	dsn := GetPostgresDsn(host, user, password, port, database, sslmode)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	dbCOntext, _ := context.
		WithTimeout(
			context.Background(),
			30*time.Second,
		)
	if err = conn.PingContext(dbCOntext); err != nil {
		panic(err)
	}
	return conn, nil
}
