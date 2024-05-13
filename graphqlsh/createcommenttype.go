package graphqlsh

import "github.com/graphql-go/graphql"

func CreateCommentType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"pid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"uid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"peid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
}
