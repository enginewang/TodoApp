package model

import "github.com/globalsign/mgo/bson"

type SimpleUser struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string        `bson:"name" json:"name"`
	Score    int           `bson:"score" json:"score"`
}

func (u *SimpleUser) GenerateID() {
	if u.Id == "" {
		u.Id = bson.NewObjectId()
	}
}
