package postgresdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connecting to the database
func InitDB() (*sql.DB, error) {
	var err error
	connStr := "host=localhost port=5432 user=admin password=12345 dbname=db_post_comment sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Нет подключения к БД")
	}
	return db, nil
}
