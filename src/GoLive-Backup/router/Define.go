package router

import (
  "gopkg.in/mgo.v2/bson"
)

type LiveMessage struct {
  Method string `json:"method" bson:"method"`
  Type string `json:"type" bson:"type"`
  Status int `json:"status" bson:"status"`
  Data bson.M `json:"data" bson:"data"`
}

type ResponseMessage struct {
  Method string `json:"method" bson:"method"`
  Type string `json:"type" bson:"type"`
  Status int `json:"status" bson:"status"`
  Message string `json:"message" bson:"message"`
}

type NetworkData struct {
  Network bson.M `json:"network" bson:"network"`
  IceCandidate bson.M `json:"candidate" bson:"candidate"`
}
