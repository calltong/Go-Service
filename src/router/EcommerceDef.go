package router

import (
  "gopkg.in/mgo.v2/bson"
  "ecommerce"
)

type LazadaProductItem struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Status string `json:"status" bson:"status"`
  Content ecommerce.LazadaRequest `json:"content" bson:"content"`
}

type StreetProductItem struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Status string `json:"status" bson:"status"`
  ProductId string `json:"product_id" bson:"product_id"`
  Content ecommerce.StreetProductRequest `json:"content" bson:"content"`
}
