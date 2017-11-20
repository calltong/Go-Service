package router

import (
  "net/http"
  "gopkg.in/mgo.v2/bson"
)

type Total struct {
	Value int `json:"value"`
}

type ResultMessage struct {
	Result string `json:"result"`
}

type Data struct {
	Data bson.M `json:"data" bson:",inline"`
}

type DataList []Data

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
  RequiredToken bool
}

type RouteList []Route
