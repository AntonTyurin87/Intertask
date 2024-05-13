package graphqlsh

import (
	//"intertask/graphqlsh"
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
)

// func QueryType(postType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
func QueryType(storage postgresdb.Storage) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "BlogQuery",
			Fields: graphql.Fields{
				"posts": &graphql.Field{
					Type: graphql.NewList(PostType),
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
						return postgresdb.AllPosts(&storage, limit, offset)
					},
				},
				"post": &graphql.Field{
					Type: PostType,
					//Type: CreatePostType(storage),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						id := p.Args["id"]
						v, _ := id.(int)
						// Read limit
						//log.Printf("fetching post with id: %d", v)

						return postgresdb.PostById(&storage, v)
					},
				},
			},
		},
	)
}
