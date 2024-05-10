package main

import (
	"net/http"

	handler "intertask/handler"
	postgresdb "intertask/postgresdb"

	_ "github.com/lib/pq"
)

func main() {

	db, _ := postgresdb.InitDB() // надо бы ошибку обработать

	storage := postgresdb.NewStorage(db)

	handler := handler.HandlerPosts(*storage)

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil) //127.0.0.1:8080
}
