package main

import (
	"database/sql"
	"net/http"

	hendlerdb "intertask/hendlerdb"
	postgresdb "intertask/postgresdb/workwithdb"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	db, _ = postgresdb.InitDB() // надо бы ошибку обработать

	handler := hendlerdb.HandlerPosts(db)

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil) //127.0.0.1:8080
}
