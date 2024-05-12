package createtype

import (
	"log"

	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
	//postgresdb "intertask/postgresdb"
)

func QueryTypeOnePost(postType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
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

						return postgresdb.PostById(&storage, v)
					},
				},
			},
		},
	)
}
