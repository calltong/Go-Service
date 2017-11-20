package common

import (
  "gopkg.in/mgo.v2/bson"
)

type NetworkData struct {
  Network bson.M `json:"network" bson:"network"`
  IceCandidate bson.M `json:"candidate" bson:"candidate"`
}

type LoginData struct {
  Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
