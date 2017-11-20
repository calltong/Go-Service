package admin

import (
  "gopkg.in/mgo.v2/bson"
  "GoLive/common"
)

// DB
type AdminData struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  GroupCode string `json:"group_code" bson:"group_code"`
  Company string `json:"company" bson:"company"`
	Name string `json:"name" bson:"name"`
  Image string `json:"image" bson:"image"`
  Type string `json:"type" bson:"type"`
  Maximun int `json:"maximun" bson:"maximun"`
  Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type ActiveAdminData struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  GroupCode string `json:"group_code" bson:"group_code"`
  Company string `json:"company" bson:"company"`
	Name string `json:"name" bson:"name"`
  Image string `json:"image" bson:"image"`
  Maximun int `json:"maximun" bson:"maximun"`
  Counter int `json:"counter" bson:"counter"`
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
  NetUser common.NetworkData `json:"net_user" bson:"net_user"`
}

// Answer network from Admin
type UserNotifyRes struct {
  SessionId string `json:"session_id" bson:"session_id"`
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  NetAdmin common.NetworkData `json:"net_admin" bson:"net_admin"`
}
