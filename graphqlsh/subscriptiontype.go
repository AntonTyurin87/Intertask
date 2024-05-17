package graphqlsh

import (
	"github.com/graphql-go/graphql"
)

func SubscriptionType(storage Blog) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Subscription",
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

					return ch, nil
				},
			},
		},
	})
}
