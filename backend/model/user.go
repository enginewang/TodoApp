package model

import "github.com/globalsign/mgo/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	Avatar   string        `bson:"avatar" json:"avatar"`
	Type     int           `bson:"type" json:"type"`
	Name     string        `bson:"name" json:"name"`
	Score    int           `bson:"score" json:"score"`
	Password string        `bson:"password" json:"password"`
	Todo     string        `bson:"todo" json:"todo"`
}

type NewTodo struct {
	Todo     string        `bson:"todo" json:"todo"`
}

func (u *User) GenerateID() {
	if u.Id == "" {
		u.Id = bson.NewObjectId()
	}
}
