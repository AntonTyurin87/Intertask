package main

import (
	"log"
	"net/http"

	"intertask/graphqlsh"
	"intertask/handler"
	"intertask/inmemory"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {

	//var err error

	//db, _ := postgresdb.InitDB()

	var InMe []inmemory.InMemoryType

	//storage := postgresdb.NewStorage(db)
	storage := inmemory.NewInMemory(InMe)

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        graphqlsh.QueryType(storage),
			Mutation:     graphqlsh.MutationType(storage),
			Subscription: graphqlsh.SubscriptionType(storage),
		})

	handler3 := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Playground: true,
	})

	http.Handle("/subscriptions", handler.NewSubscriptionHandler(schema))

	http.Handle("/graphql", handler3)
	log.Println("Server started at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
