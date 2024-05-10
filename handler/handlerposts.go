package handler

import (
	postgresdb "intertask/postgresdb"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func HandlerPosts(storage postgresdb.Storage) *handler.Handler {
	postType := CreatePostType()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: QueryType(postType, storage),
		},
	)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	return handler
}
