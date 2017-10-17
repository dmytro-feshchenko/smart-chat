package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	// User - user account information
	User struct {
		ID       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name     string        `json:"name"`
		Email    string        `json:"email"`
		Password string        `json:"password"`
	}
	// UserCredentials - contains fields for authentication
	UserCredentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	// Token - represents JWT token for authorization
	Token struct {
		Token string `json:"token"`
	}
)
