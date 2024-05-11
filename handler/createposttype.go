package handler

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"

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
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					post, _ := p.Source.(*postgresdb.Post)
					log.Printf("fetching comments of post with id: %d", post.PID)
					fmt.Println("1")
					return storage.FetchCommentsByPostID(post.PID)
				},
			},
		},
	})
}
