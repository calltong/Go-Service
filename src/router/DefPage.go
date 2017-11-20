package router

import (
  "gopkg.in/mgo.v2/bson"
)

type Page struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Page string `json:"page" bson:"page"`
	Name string `json:"name" bson:"name"`
	Update int `json:"updated" bson:"updated"`
	Status string `json:"status" bson:"status"`
  Data bson.M `json:"data" bson:"data"`
}

type PageList []Page

type CssData struct {
  Size float32 `json:"size" bson:"size"`
	BgColor string `json:"bg_color" bson:"bg_color"`
  Color string `json:"color" bson:"color"`
	Font string `json:"font" bson:"font"`
}

type HomeData struct {
  Css CssData `json:"css" bson:"css"`
  List []HomeContent `json:"list" bson:"list"`
}

type AboutusData struct {
  Title string `json:"title" bson:"title"`
  Description string `json:"description" bson:"description"`
  List []ContentItem `json:"list" bson:"list"`
}

/// Content Defination.

type HomeContent struct {
	Type string `json:"type" bson:"type"`
  Data ContentData `json:"data" bson:"data"`
}

type ContentData struct {
	Title string `json:"title" bson:"title"`
  Description string `json:"description" bson:"description"`
  List []ContentItem `json:"list" bson:"list"`
}

type ContentItem struct {
	Type string `json:"type" bson:"type"`
  Value string `json:"value" bson:"value"`
  Title string `json:"title" bson:"title"`
  Name string `json:"name" bson:"name"`
  Preview string `json:"preview" bson:"preview"`
}
