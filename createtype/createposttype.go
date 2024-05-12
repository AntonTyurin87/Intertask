package createtype

import (
	"log"

	"github.com/graphql-go/graphql"

	blogInterface "intertask/cmd/bloginterface"
	postgresdb "intertask/postgresdb"
)

func CreatePostType(commentType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"postauthorid": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"cancomment": &graphql.Field{
				Type: graphql.Boolean,
			},
			"comments": &graphql.Field{
				Type: graphql.NewList(commentType),
				Args: graphql.FieldConfigArgument{
					"limit": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"offset": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post, _ := p.Source.(*blogInterface.Post)
					log.Printf("fetching comments of post with id: %d", post.PID)
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

					return postgresdb.CommentsByPostID(&storage, post.PID, limit, offset)
				},
			},
		},
	})
}
