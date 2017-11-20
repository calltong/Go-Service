package router

import (
  "gopkg.in/mgo.v2/bson"
)

type SocialData struct {
  Type string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`
	Id string `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
	Image string `json:"image" bson:"image"`
}

type CustomerAddress struct {
  Address string `json:"address" bson:"address"`
	City string `json:"city" bson:"city"`
	Postcode string `json:"postcode" bson:"postcode"`
	Mobile string `json:"mobile" bson:"mobile"`
}

type Customer struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Browser string `json:"browser" bson:"browser`
  Type string `json:"type" bson:"type"`
	Name string `json:"name" bson:"name"`
  Email string `json:"email" bson:"email"`
  Password string `json:"password" bson:"password"`
	Status string `json:"status" bson:"status"`
	Information CustomerAddress `json:"information" bson:"information"`
  MediaList []SocialData `json:"media_list" bson:"media_list"`
	LastUpdated int64 `json:"last_updated" bson:"last_updated"`
}

type CustomerList []Customer
