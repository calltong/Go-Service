package router

import (
  "gopkg.in/mgo.v2/bson"
)

type LogValue struct {
  Type string `json:"type" bson:"type"`
  Date int64 `json:"date" bson:"date"`
  Message string `json:"message" bson:"message"`
  Source string `json:"source" bson:"source"`
}

type Logger struct {
  Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Year string `json:"year" bson:"year"`
  System string `json:"system" bson:"system"`
  List []LogValue `json:"list" bson:"list"`
}
