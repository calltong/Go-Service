package user

import (
  "gopkg.in/mgo.v2/bson"
  "GoLive/common"
)

type ResUserId struct {
  UserId bson.ObjectId `json:"user_id" bson:"user_id"`
}

type PersonalData struct {
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Mobile string `json:"mobile" bson:"mobile"`
}

type UserData struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Personal PersonalData `json:"personal" bson:"personal"`
	Browser bson.M `json:"browser" bson:"browser"`
}

type ReqConnect struct {
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  NetUser common.NetworkData `json:"net_user" bson:"net_user"`
}

type ResConnect struct {
  SessionId string `json:"session_id" bson:"session_id"`
  AdminId string `json:"admin_id" bson:"admin_id"`
  NetAdmin common.NetworkData `json:"net_admin" bson:"net_admin"`
}

type ReqReconnect struct {
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  SessionId string `json:"session_id" bson:"session_id"`
  AdminId string `json:"admin_id" bson:"admin_id"`
  NetUser common.NetworkData `json:"net_user" bson:"net_user"`
}
