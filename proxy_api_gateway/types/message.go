package types

import "github.com/graphql-go/graphql"

// MessageType - GraphQL ObjectType which represents message in chat
var MessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Message",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: UserType,
		},
	},
})
