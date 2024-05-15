package main

import (
	"encoding/json"
	"fmt"
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

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {

	db, _ := postgresdb.InitDB()

	storage := postgresdb.NewStorage(db)

	schema, _ := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        graphqlsh.QueryType(storage),
			Mutation:     graphqlsh.MutationType(storage),
			Subscription: graphqlsh.SubscriptionType(storage),
		})

	//handler3 := Handler(storage)

	handler3 := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})

	http.HandleFunc("/subscription", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	http.Handle("/graphql", handler3)
	log.Println("Server started at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/////////////////////////////////////////////////////
/*
func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query,
	})
	if len(result.Errors) > 0 {
			fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}


	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

///////////////////////////////////////////////////////////

func Handler(storage graphqlsh.Blog) *gqlhandler.Handler {

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:        graphqlsh.QueryType(storage),
			Mutation:     graphqlsh.MutationType(storage),
			Subscription: graphqlsh.SubscriptionType(storage),
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
*/
