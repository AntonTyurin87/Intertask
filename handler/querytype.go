package handler

import (
	postgresdb "intertask/postgresdb"

	"github.com/graphql-go/graphql"

	_ "github.com/lib/pq"
)

func QueryType(postType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"Posts": &graphql.Field{
					Type: graphql.NewList(postType),
					Args: graphql.FieldConfigArgument{
						"limit": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
						"offset": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// Read limit
						limit, _ := p.Args["limit"].(int)
						if limit <= 0 || limit > 20 {
							limit = 10
						}
						// Read offset
						offset, _ := p.Args["offset"].(int)
						if offset < 0 {
							offset = 0
						}

						//fmt.Println(p.)

						return storage.GetPosts(limit, offset)
					},
				},
			},
		},
	)
}
