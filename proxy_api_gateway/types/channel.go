package types

import "github.com/graphql-go/graphql"

// ChannelType - GraphQL ObjectType which represents chat channel
var ChannelType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Channel",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"channels": &graphql.Field{
			Type: graphql.NewList(MessageType),
		},
	},
})
