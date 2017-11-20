package router

import (
  "gopkg.in/mgo.v2/bson"
)

type LazadaType struct {
  CategoryId int `json:"category_id" bson:"category_id"`
  Model string `json:"model" bson:"model"`
}

type EcommerceType struct {
	Lazada LazadaType `json:"lazada" bson:"lazada"`
}

type ProductTypeLangEng struct {
	Name string `json:"name" bson:"name"`
}

type InfoType struct {
	Name string `json:"name" bson:"name"`
}

type ContentType struct {
	Main InfoType `json:"main" bson:"main"`
  English InfoType `json:"english" bson:"english"`
}


type ProductType struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Content ContentType `json:"content" bson:"content"`
  Ecommerce EcommerceType `json:"ecommerce" bson:"ecommerce"`
}

type ProductTypeList []ProductType
