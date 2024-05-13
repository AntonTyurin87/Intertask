package graphqlsh

import (
	blogInterface "intertask/cmd/bloginterface"
	"intertask/postgresdb"

	"github.com/graphql-go/graphql"
)

// Нужно сперва вернуть UserID

func MutationType(storage postgresdb.Storage) *graphql.Object {
	//func MutationType(storage postgresdb.Storage) *graphql.Object {

	/*
		postType1 := graphql.NewObject(graphql.ObjectConfig{
			Name: "Post1",
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
			},
		},
		)
	*/

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BlogMutation",
		Fields: graphql.Fields{
			"createpost": &graphql.Field{
				Type: PostType,
				//Type: CreatePostType(storage),
				//Type: graphql.NewList(CreatePostType(storage)),
				//Description: "Create new Post",
				Args: graphql.FieldConfigArgument{
					"postauthorid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"cancomment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					postauthorid, _ := params.Args["postauthorid"].(int)
					text, _ := params.Args["text"].(string)
					cancomment, _ := params.Args["cancomment"].(bool)

					newPost := blogInterface.Post{
						PostAuthorID: postauthorid,
						Text:         text,
						CanComment:   cancomment,
					}

					//fmt.Println("1")
					return postgresdb.CreatePost(&storage, &newPost)
				},
			},

			//"CommentStatus":
		},
	})
}
