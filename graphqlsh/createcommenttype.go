package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

/*
	func CreateCommentType(storage postgresdb.Storage) *graphql.Object {
		return graphql.NewObject(graphql.ObjectConfig{
			Name: "Comment",
			Fields: graphql.Fields{
				"comments": &graphql.Field{
					Type: graphql.NewList(CommentType),
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
						return
					},
				},
			},
		})
	}
*/
func CreateCommentType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Comment2",
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
