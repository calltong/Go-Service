package router

import (
  "gopkg.in/mgo.v2/bson"
)

type GoogleAnalyticData struct {
	Id string `json:"id" bson:"id"`
}

type LazadaUser struct {
	Url string `json:"url" bson:"url"`
  Id string `json:"user_id" bson:"user_id"`
  Version string `json:"version" bson:"version"`
  Key string `json:"key" bson:"key"`
}

type LazadaContent struct {
	Brand string `json:"brand" bson:"brand"`
  WarrantyType string `json:"warranty_type" bson:"warranty_type"`
  Warranty string `json:"warranty" bson:"warranty"`
  PackageLength string `json:"pkage_length" bson:"package_length"`
  PackageWeight string `json:"package_weight" bson:"package_weight"`
  PackageHeight string `json:"package_height" bson:"package_height"`
  PackageWidth string `json:"package_width" bson:"package_width"`
  BraType string `json:"bras_types" bson:"bras_types"`
  SleeperType string `json:"sleep_lounge_styles" bson:"sleep_lounge_styles"`
  DurationPromotion int `json:"duration_promotion" bson:"duration_promotion"`
}

type LazadaData struct {
	User LazadaUser `json:"user" bson:"user"`
  Product LazadaContent `json:"product" bson:"product"`
}

type StreetUser struct {
	Url string `json:"url" bson:"url"`
  Version string `json:"version" bson:"version"`
  Key string `json:"key" bson:"key"`
}

type StreetContent struct {
	Brand string `json:"brand" bson:"brand"`
  Warranty string `json:"warranty_type" bson:"warranty_type"`
  PackageLength string `json:"pkage_length" bson:"package_length"`
  PackageWeight string `json:"package_weight" bson:"package_weight"`
  PackageHeight string `json:"package_height" bson:"package_height"`
  PackageWidth string `json:"package_width" bson:"package_width"`
  BraType string `json:"bras_types" bson:"bras_types"`
  SleeperType string `json:"sleep_lounge_styles" bson:"sleep_lounge_styles"`
  DurationPromotion int `json:"duration_promotion" bson:"duration_promotion"`
}

type StreetData struct {
	User StreetUser `json:"user" bson:"user"`
  Product StreetContent `json:"product" bson:"product"`
}

type SettingStore struct {
	ID bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Type string `json:"type" bson:"type"`
  GoogleAnalytic GoogleAnalyticData `json:"google_analytic" bson:"google_analytic"`
  Lazada LazadaData `json:"lazada" bson:"lazada"`
  Street StreetData `json:"street" bson:"street"`
}
