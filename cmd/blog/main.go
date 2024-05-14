package main

import (
	"log"
	"net/http"

	"intertask/graphqlsh"
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	//_ "github.com/lib/pq"
	//handler "intertask/handler"
	// postgresdb "intertask/postgresdb"
)

func main() {

	db, _ := postgresdb.InitDB()

	storage := postgresdb.NewStorage(db)

	handler3 := Handler(storage)

	http.Handle("/graphql", handler3)
	log.Println("Server started at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(storage graphqlsh.Blog) *gqlhandler.Handler {

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    graphqlsh.QueryType(storage),
			Mutation: graphqlsh.MutationType(storage),
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})

	return handler
}
