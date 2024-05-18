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

// The function that creates a type for sending mutations to the database
func MutationType(storage Blog) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "BlogMutation",
		Description: "A type for making changes to the database of posts and comments and to the posts themselves.",
		Fields: graphql.Fields{
			"createpost":    createPost(storage),
			"commentstatus": commentStatus(storage),
			"createcomment": createComment(storage),
		},
	})
}

// The function that creates a type for sending requests to the database to create a post.
func createPost(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type:        PostType,
		Description: "Type for creating a new post based on the base type.",
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

// The function that creates a type for sending requests to the database to change the comment status of a post.
func commentStatus(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type:        PostType,
		Description: "Type for making changes to an existing post, based on the base type.",
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

// The function that creates a type for sending requests to the database to create a comment on a post.
func createComment(storage Blog) *graphql.Field {
	return &graphql.Field{
		Type:        CommentType,
		Description: "Type for creating a new comment based on the base type.",
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

			// Intercept information about a new comment and send it to work for subscription.
			NewComment(postid, newComment)
			return storage.CreateNewComment(&newComment)
		},
	}
}
