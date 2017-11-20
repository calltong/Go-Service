package router

import (
  "gopkg.in/mgo.v2/bson"
)

type LazadaSize struct {
	Code string `json:"code" bson:"code"`
}

type EcommerceSize struct {
	Lazada LazadaSize `json:"lazada" bson:"lazada"`
}

type InfoSize struct {
	Name string `json:"name" bson:"name"`
}

type ContentSize struct {
	Main InfoType `json:"main" bson:"main"`
  English InfoType `json:"english" bson:"english"`
}

type ProductSize struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Code string `json:"code" bson:"code"`
  Content ContentSize `json:"content" bson:"content"`
  Ecommerce EcommerceSize `json:"ecommerce" bson:"ecommerce"`
}

type ProductSizeList []ProductSize
