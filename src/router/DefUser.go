package router

import (
	"gopkg.in/mgo.v2/bson"
)
type TokenData struct {
	Token string `json:"token"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type UserProject struct {
	Id bson.ObjectId `json:"id" bson:"id,omitempty"`
	Role string `json:"role" bson:"role"`
}

type User struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Mobile string `json:"mobile" bson:"mobile"`
	ProjectList []UserProject `json:"project_list" bson:"project_list"`
}
