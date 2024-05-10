package hendlerdb

import "github.com/graphql-go/graphql"

// Создание объекта GraphQL для описания постов
func CreatePostType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Post",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"ptext": &graphql.Field{
					Type: graphql.String,
				},
				"uid": &graphql.Field{
					Type: graphql.Int,
				},
				"cancomment": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		},
	)
}
