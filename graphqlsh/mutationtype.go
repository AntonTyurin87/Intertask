package graphqlsh

import (
	"fmt"
	blogInterface "intertask/cmd/bloginterface"

	"github.com/graphql-go/graphql"
)

// Нужно сперва вернуть UserID

func MutationType(storage Blog) *graphql.Object {
	//func MutationType(storage postgresdb.Storage) *graphql.Object {

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
					return storage.CreateNewPost(&newPost)
				},
			},
			"commentstatus": &graphql.Field{
				Type: PostType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"postauthorid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"cancomment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					postauthorid, _ := params.Args["postauthorid"].(int)
					cancomment, _ := params.Args["cancomment"].(bool)

					correctPost := blogInterface.Post{
						PID:          id,
						PostAuthorID: postauthorid,
						CanComment:   cancomment,
					}
					return storage.CorrectPost(&correctPost)
				},
			},
			"createcomment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"pid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"uid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"peid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					text, _ := params.Args["text"].(string)
					uid := params.Args["uid"].(int)
					pid, _ := params.Args["pid"].(int)
					peid, err := params.Args["peid"].(int)
					if !err {
						fmt.Println("1")
						peid = 0
					}

					newComment := blogInterface.Comment{
						UserID:   uid,
						PostID:   pid,
						PerentID: peid,
						Text:     text,
					}
					return storage.CreateNewComment(&newComment)
				},
			},
			"dosubscription": &graphql.Field{
				Type: UserSubscription,
				Args: graphql.FieldConfigArgument{
					"uid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"pid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"confirmation": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					confirmation, _ := params.Args["confirmation"].(bool)
					uid, _ := params.Args["uid"].(int)
					pid, _ := params.Args["pid"].(int)

					newSubscription := blogInterface.UserSubscription{
						UserID:       uid,
						PostID:       pid,
						Сonfirmation: confirmation,
					}
					return storage.CreateUserSubscription(&newSubscription)
				},
			},
		},
	})
}
