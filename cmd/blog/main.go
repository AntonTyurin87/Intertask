package main

import (
	"log"
	"net/http"

	"intertask/handler"
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	//_ "github.com/lib/pq"
	//handler "intertask/handler"
	// postgresdb "intertask/postgresdb"
)

type PostById interface {
	GetPostById()
}

func GetPost(something PostById) {
	something.GetPostById()
}

func main() {

	db, _ := postgresdb.InitDB()

	storage := postgresdb.NewStorage(db)

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(
			handler.CreateQueryType(handler.CreatePostType(handler.CreateCommentType(), *storage), *storage),
		),
	})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
	})

	http.Handle("/graphql", handler)
	log.Println("Server started at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func main() {

	db, _ := postgresdb.InitDB() // надо бы ошибку обработать

	storage := postgresdb.NewStorage(db)

	handler := handler.HandlerPosts(*storage)

	http.Handle("/graphql", handler)
	http.ListenAndServe(":8080", nil) //127.0.0.1:8080
}
*/
