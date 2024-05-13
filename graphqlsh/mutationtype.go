package graphqlsh

import (
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
)

// Нужно сперва вернуть UserID

func MutationType(postType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BlogMutation",
		Fields: graphql.Fields{
			"createPost": &graphql.Field{
				Type: postType,
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"cancomment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
			},
		},
	})
}
