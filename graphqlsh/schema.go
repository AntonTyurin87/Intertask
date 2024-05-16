package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

type Post struct {
	ID         int    `json:"id"`
	Text       string `json:"text"`
	UserID     int    `json:"usreid"`
	CanComment bool   `json:"cancomment"`
}

type Comment struct {
	ID       int    `json:"id"`
	UserID   int    `json:"userid"`
	Text     string `json:"text"`
	PostID   int    `json:"postid"`
	PerentID int    `json:"perentid"`
}

type UserSubscription struct {
	ID           int  `json:"id"`
	UserID       int  `json:"userid"`
	PostID       int  `json:"postid"`
	Сonfirmation bool `json:"confirmation"`
}

var PostType = graphql.NewObject(graphql.ObjectConfig{
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
			Type: CommentType,
		},
	},
},
)

var CommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
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
			Type: graphql.String,
		},
	},
})

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
