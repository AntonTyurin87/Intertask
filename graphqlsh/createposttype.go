package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

func CreatePostType(storage Blog) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"userid": &graphql.Field{
				Type: graphql.Int,
			},
			"cancomment": &graphql.Field{
				Type: graphql.Boolean,
			},
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
					post, _ := p.Source.(*Post)
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

					return storage.FetchCommentsByPostID(post.ID, limit, offset)

				},
			},
		},
	},
	)
}
