package notification

import (
  "gopkg.in/mgo.v2/bson"
)

type NotifyMessage struct {
  Message bson.M `json:"message" bson:"message"`
}
