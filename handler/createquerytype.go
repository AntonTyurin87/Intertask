package handler

import (
	"log"

	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
	//postgresdb "intertask/postgresdb"
)

func CreateQueryType(postType *graphql.Object, storage postgresdb.Storage) graphql.ObjectConfig {
	return graphql.ObjectConfig{Name: "QueryType", Fields: graphql.Fields{
		"post": &graphql.Field{
			Type: postType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"]
				v, _ := id.(int)
				log.Printf("fetching post with id: %d", v)
				return storage.FetchPostByiD(v)
			},
		},
	}}
}
