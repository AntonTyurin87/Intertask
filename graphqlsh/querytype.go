package graphqlsh

import (
	//"intertask/graphqlsh"

	"github.com/graphql-go/graphql"

	blogInterface "intertask/cmd/bloginterface"
)

// Blog Interface
type Blog interface {
	FetchAllPosts(limit, offset int) ([]blogInterface.Post, error)
	FetchPostByiD(id int) (*blogInterface.Post, error)
	FetchCommentsByPostID(id, limit, offset int) ([]blogInterface.Comment, error)
	CreateNewPost(newPost *blogInterface.Post) (*blogInterface.Post, error)
	CreateNewComment(newComment *blogInterface.Comment) (*blogInterface.Comment, error)
	CorrectPost(correctPost *blogInterface.Post) (*blogInterface.Post, error)
	CreateUserSubscription(newSubscription *blogInterface.UserSubscription) (*blogInterface.UserSubscription, error)
	CreateNotification(comment int) ([]blogInterface.UserSubscription, error)
}

// func QueryType(postType *graphql.Object, storage postgresdb.Storage) *graphql.Object {
func QueryType(storage Blog) *graphql.Object {
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
						return storage.FetchAllPosts(limit, offset)
					},
				},
				"post": &graphql.Field{
					Type: CreatePostType(storage),
					//Type: PostType,
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

						return storage.FetchPostByiD(v)
					},
				},
			},
		},
	)
}
