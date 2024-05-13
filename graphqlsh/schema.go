package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

var PostType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Post",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"postauthorid": &graphql.Field{
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
