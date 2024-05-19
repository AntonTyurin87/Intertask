package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"intertask/graphqlsh"
	"intertask/handler"
	"intertask/inmemory"
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {

	var storage graphqlsh.Blog

	//Set in Docker settings
	boolValue := os.Getenv("IN_MEMORY")    //Operating system variable for selecting memory or database mode.
	urlAdress := os.Getenv("POSTGRES_URL") //Operating system variable for generating a URL for connecting to the database.

	if boolValue == "true" {
		var InMe []inmemory.InMemoryType
		storage = inmemory.NewInMemory(InMe)
	} else {
		db, err := postgresdb.InitDB(urlAdress)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		storage = postgresdb.NewStorage(db)
	}

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        graphqlsh.QueryType(storage),
			Mutation:     graphqlsh.MutationType(storage),
			Subscription: graphqlsh.SubscriptionType(storage),
		})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	handler3 := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Playground: true,
	})

	http.Handle("/subscriptions", handler.NewSubscriptionHandler(schema))

	http.Handle("/graphql", handler3)
	log.Println("Server started at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
