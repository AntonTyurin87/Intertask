package graphqlsh

import (
	"slices"
	"sync"

	"github.com/graphql-go/graphql"
)

var (
	mx                   sync.Mutex
	commentSubscriptions = map[int][]chan any{}
)

func SubscribeToNewComments(postID int, ch chan any) {
	mx.Lock()

	commentSubscriptions[postID] = append(commentSubscriptions[postID], ch)

	mx.Unlock()
}

func UnsubscribeFromNewComments(postID int, ch chan any) {
	mx.Lock()

	idx := slices.Index(commentSubscriptions[postID], ch)
	commentSubscriptions[postID] = slices.Delete(commentSubscriptions[postID], idx, idx+1)

	mx.Unlock()
}

func NewComment(postID int, newComment any) {
	mx.Lock()

	for _, subscription := range commentSubscriptions[postID] {
		subscription <- newComment
	}

	mx.Unlock()
}

func MutationType(storage Blog) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "BlogMutation",
		Fields: graphql.Fields{
			"createpost":    createPost(storage),
			"commentstatus": commentStatus(storage),
			"createcomment": createComment(storage),
		},
	})
}

func createPost(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type: PostType,
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
	}
}

func commentStatus(storage Blog) *graphql.Field {
	return &graphql.Field{
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
	}
}

func createComment(storage Blog) *graphql.Field {
	return &graphql.Field{
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
	}
}
