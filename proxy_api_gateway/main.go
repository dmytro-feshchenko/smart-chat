package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"net/http"

	"github.com/graphql-go/graphql"
	types "github.com/technoboom/smart-chat/proxy_api_gateway/types"
)

type (
	// Message - contains data about single chat message
	Message struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	}
)

var MessageList []Message

func init() {
	message1 := Message{ID: 1, Text: "Message 1"}
	message2 := Message{ID: 2, Text: "Message 2"}
	message3 := Message{ID: 3, Text: "Message 3"}

	MessageList = append(MessageList, message1, message2, message3)

	rand.Seed(time.Now().UnixNano())
}

// rootMutation - root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"registerAccount": &graphql.Field{
			Type:        types.UserType,
			Description: "register a new account",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// name, _ := params.Args["name"].(string)
				// email, _ := params.Args["email"].(string)

				// newAccount := User{
				// 	Name:  name,
				// 	Email: email,
				// }

				return nil, nil
			},
		},
	},
})

// type Channel struct {
// 	Name: string
// 	Channels: []
// }

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user": &graphql.Field{
			Type:        types.UserType,
			Description: "Get single user details",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, isOk := params.Args["id"].(string)
				if isOk {
					// todo: find and return user info
				}

				return nil, nil
			},
		},
		"usersList": &graphql.Field{
			Type:        graphql.NewList(types.UserType),
			Description: "Get list of users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			},
		},
		"channels": &graphql.Field{
			Type:        graphql.NewList(types.ChannelType),
			Description: "Get list of channels",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				data, _ := http.NewRequest("GET", "localhost:6464/channels", nil)
				return data, nil
			},
		},
		"messages": &graphql.Field{
			Type:        graphql.NewList(types.MessageType),
			Description: "Get messages list",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return MessageList, nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

// Entry point for the program
func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
