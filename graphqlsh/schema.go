package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

// Data structure for posts.
type Post struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	UserID     int    `json:"usreid"`
	CanComment bool   `json:"cancomment"`
}

// Data structure for comments.
type Comment struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userid"`
	Text     string `json:"text"`
	PostID   int    `json:"postid"`
	PerentID int    `json:"perentid"`
}

// Basic type for posts.
var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post_Base",
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
			Type: CommentType,
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"offset": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
		},
	},
},
)

// Basic type for comments.
var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment_Base",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"postid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"userid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"perentid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"text": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

/*
// Basic type for subscriptions. Probably not needed.
var UserSubscriptionType = graphql.NewObject(graphql.ObjectConfig{
	Name: "UserSubscription",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"userid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"postid": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"confirmation": &graphql.Field{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
	},
},
)
*/
