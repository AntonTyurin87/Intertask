package graphqlsh

import (
	"fmt"

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
					"userid": &graphql.ArgumentConfig{
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

					userid, _ := params.Args["userid"].(int)
					text, _ := params.Args["text"].(string)
					cancomment, _ := params.Args["cancomment"].(bool)

					newPost := Post{
						UserID:     userid,
						Text:       text,
						CanComment: cancomment,
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
					"userid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"cancomment": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					id, _ := params.Args["id"].(int)
					userid, _ := params.Args["userid"].(int)
					cancomment, _ := params.Args["cancomment"].(bool)

					correctPost := Post{
						ID:         id,
						UserID:     userid,
						CanComment: cancomment,
					}
					return storage.UpdatePost(&correctPost)
				},
			},
			"createcomment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"postid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"userid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"perentid": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
					"text": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					text, _ := params.Args["text"].(string)
					userid := params.Args["userid"].(int)
					postid, _ := params.Args["postid"].(int)
					perentid, err := params.Args["perentid"].(int)

					if !err {
						fmt.Println("1")
						perentid = 0
					}

					newComment := Comment{
						UserID:   userid,
						PostID:   postid,
						PerentID: perentid,
						Text:     text,
					}
					return storage.CreateNewComment(&newComment)
				},
			},
			"dosubscription": &graphql.Field{
				Type: UserSubscriptionType,
				Args: graphql.FieldConfigArgument{
					"userid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"postid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"confirmation": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Boolean),
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {

					confirmation, _ := params.Args["confirmation"].(bool)
					userid, _ := params.Args["userid"].(int)
					postid, _ := params.Args["postid"].(int)

					newSubscription := UserSubscription{
						UserID:       userid,
						PostID:       postid,
						Сonfirmation: confirmation,
					}
					return storage.CreateUserSubscription(&newSubscription)
				},
			},
		},
	})
}
