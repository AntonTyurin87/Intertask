package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

// The function that creates a type for sending subscription information to the go channel.
func SubscriptionType(storage Blog) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Subscription",
		Description: "Type for sending subscription information to the go channel.",
		Fields: graphql.Fields{
			"createcomment": &graphql.Field{
				Type: CommentType,
				Args: graphql.FieldConfigArgument{
					"postid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source, nil
				},
				// Subscribe and create a separate go channel for it.
				Subscribe: func(params graphql.ResolveParams) (interface{}, error) {
					ch := make(chan any)
					postid, _ := params.Args["postid"].(int)

					SubscribeToNewComments(postid, ch)
					go func() {
						for {
							select {
							case <-params.Context.Done():
								UnsubscribeFromNewComments(postid, ch)
								close(ch)
								return
							}
						}
					}()
					// Returns the go channel.
					return ch, nil
				},
			},
		},
	})
}
