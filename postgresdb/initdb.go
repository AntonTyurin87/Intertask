package postgresdb

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Create connecting to the database
func InitDB(urlAdress string) (*sql.DB, error) {

	var err error

	connStr := urlAdress //"host=localhost port=5432 user=admin password=12345 dbname=db_post_comment sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("No connection to database", err)
	}
	return db, nil
}
