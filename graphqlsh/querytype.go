package graphqlsh

import (
	//"intertask/graphqlsh"

	"github.com/graphql-go/graphql"
)

// Blog Interface for workштп with the data storage.
type Blog interface {
	//Gets all posts from data storage.
	FetchAllPosts(limit, offset int) ([]Post, error)

	// Get a post and comments to it by ID from data storage.
	FetchPostByiD(id int) (*Post, error)

	//Get comments for a specific post from data storage.
	FetchCommentsByPostID(id, limit, offset int) ([]Comment, error)

	//Creates a record of a new post in data storage.
	CreateNewPost(newPost *Post) (*Post, error)

	// Creates a record of a new comment in data storage.
	CreateNewComment(newComment *Comment) (*Comment, error)

	// Makes a change to the post entry about the ability to comment the post in data storage.
	UpdatePost(correctPost *Post) (*Post, error)
}

// The function that creates a type for sending queries to the database.
func QueryType(storage Blog) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:        "BlogQuery",
			Description: "Type for queries to search for posts.",
			Fields: graphql.Fields{
				"posts": getPosts(storage),
				"post":  getPostWithComments(storage),
			},
		},
	)
}

// The function that creates a type to query the database to retrieve all posts.
func getPosts(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(PostType),
		Description: "Type for queries to find all posts based on base type.",
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
	}
}

// The function that creates a type for sending queries to the database to retrieve a specific post and all comments on it.
func getPostWithComments(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type:        CreatePostType(storage),
		Description: "Type for queries to find post by ID based on base type.",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"]
			v, _ := id.(int)

			return storage.FetchPostByiD(v)
		},
	}
}
