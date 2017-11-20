package register

import (
  "golang.org/x/net/websocket"
  "gopkg.in/mgo.v2/bson"
)

type Actor struct {
  Id string
  Group string
  Connection *websocket.Conn
}

type GroupList struct {
  List map[string]Actor
}

type LiveMessage struct {
  Method string `json:"method" bson:"method"`
  Type string `json:"type" bson:"type"`
  Status int `json:"status" bson:"status"`
  Data bson.M `json:"data" bson:"data"`
}

type ReqRegister struct {
  Group string `json:"group_code" bson:"group_code"`
  Id string `json:"id" bson:"id"`
}

// Response
type ResponseMessage struct {
  Method string `json:"method" bson:"method"`
  Type string `json:"type" bson:"type"`
  Status int `json:"status" bson:"status"`
  Message string `json:"message" bson:"message"`
}
