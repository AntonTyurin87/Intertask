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
					"pid": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					comment, _ := p.Source.(Comment)

					//Вот тут надо сделать асинхронную рассылку уведомлений!!!

					return storage.CreateNotification(comment.PostID)
				},
			},
		},
	})
}
