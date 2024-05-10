package hendlerdb

import (
	"database/sql"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func HandlerPosts(db *sql.DB) *handler.Handler {
	postType := CreatePostType()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: QueryType(postType, db),
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
