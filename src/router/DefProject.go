package router

import (
	"gopkg.in/mgo.v2/bson"
)

type ProjectGeneral struct {
	Folder string `json:"root_folder" bson:"root_folder"`
  Database string `json:"db_name" bson:"db_name"`
  Address string `json:"address" bson:"address"`
}

type ProjectFunction struct {
	Ecommerce string `json:"ecommerce" bson:"ecommerce"`
}

type ProjectFull struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Mode string `json:"mode" bson:"mode"`
	AddressList []string `json:"address_list" bson:"address_list"`
  General ProjectGeneral `json:"general" bson:"general"`
	Function ProjectFunction `json:"function" bson:"function"`
}

type Project struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Code string `json:"code" bson:"code"`
	Version string `json:"version" bson:"version"`
	Mode string `json:"mode" bson:"mode"`
	Created string `json:"created" bson:"created"`
  General ProjectGeneral `json:"general" bson:"general"`
	Function ProjectFunction `json:"function" bson:"function"`
}
