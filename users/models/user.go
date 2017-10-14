package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		ID    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `json:"name"`
		Email string        `json:"email"`
	}
)
