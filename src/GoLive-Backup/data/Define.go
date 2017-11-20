package data

import (
  "golang.org/x/net/websocket"
  "gopkg.in/mgo.v2/bson"
)

type LiveAdmin struct {
  AdminId string
  Company string
  Name string
  Max int
  Count int
  Connection *websocket.Conn
}

type adminGroup struct {
  List map[string]LiveAdmin
}

type NetworkData struct {
  Network bson.M
  IceCandidate bson.M
}

type LiveUser struct {
  Company string
  UserId string
  AdminId string
  NetUser NetworkData
  NetAdmin NetworkData
  Connection *websocket.Conn
  Done chan bool
}

type userGroup struct {
  List map[string]LiveUser
}


var mapAdmin map[string]adminGroup
var mapUser map[string]userGroup

func InitLiveData() {
  mapAdmin = make(map[string]adminGroup)
  mapUser = make(map[string]userGroup)
}


// Message
type UserNotify struct {
  Company string
  UserId string
  AdminId string
  SessionId string
  Network bson.M
  IceCandidate bson.M
  Raw []byte
}
