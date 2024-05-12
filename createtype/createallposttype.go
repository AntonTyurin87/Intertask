package createtype

import "github.com/graphql-go/graphql"

func CreateAllPostType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Post",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
				},
				"text": &graphql.Field{
					Type: graphql.String,
				},
				"postauthorid": &graphql.Field{
					Type: graphql.Int,
				},
				"cancomment": &graphql.Field{
					Type: graphql.Boolean,
				},
			},
		},
	)
}
