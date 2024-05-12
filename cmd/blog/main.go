package main

import (
	"log"
	"net/http"

	"intertask/createtype"
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

	handler1 := HandlerPostComments(storage)
	//handler2 := HandlerPosts(storage)

	http.Handle("/graphql", handler1)
	//http.Handle("/graphql", handler2)
	log.Println("Server started at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandlerPosts(storage *postgresdb.Storage) *gqlhandler.Handler {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: createtype.QueryTypePosts(createtype.CreateAllPostType(), *storage),
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

func HandlerPostComments(storage *postgresdb.Storage) *gqlhandler.Handler {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: createtype.QueryTypeOnePost(createtype.CreatePostType(createtype.CreateCommentType(), *storage), *storage),
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
