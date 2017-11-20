package router

import (
  "gopkg.in/mgo.v2/bson"
)

// DB
type AdminData struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Company string `json:"company" bson:"company"`
	Name string `json:"name" bson:"name"`
  Type string `json:"type" bson:"type"`
  Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// Message

type AdminRegisterReq struct {
  AdminId string `json:"admin_id" bson:"admin_id"`
}

// Offer network from user
type UserNotifyReq struct {
  SessionId string `json:"session_id" bson:"session_id"`
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  NetUser NetworkData `json:"net_user" bson:"net_user"`
}

// Answer network from Admin
type UserNotifyRes struct {
  SessionId string `json:"session_id" bson:"session_id"`
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  NetAdmin NetworkData `json:"net_admin" bson:"net_admin"`
}
