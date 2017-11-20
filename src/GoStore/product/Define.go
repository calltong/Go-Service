package product

import (
  "gopkg.in/mgo.v2/bson"
)

type InfoProduct struct {
  Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
  Condition string `json:"condition" bson:"condition"`
  PackageContent string `json:"package_content" bson:"package_content"`
}

type ContentProduct struct {
  Main InfoProduct `json:"main" bson:"main"`
  English InfoProduct `json:"english" bson:"english"`
}

type InfoData struct {
	Value string `json:"value" bson:"value"`
	List []string `json:"list" bson:"list"`
  PackageContent string `json:"package_content" bson:"package_content"`
}

type ImageData struct {
	Status bool `json:"status" bson:"status"`
	Data string `json:"data" bson:"data"`
}

type ImageDataList []ImageData

type Stock struct {
	Size ProductSize `json:"size" bson:"size"`
	Quantity int `json:"quantity" bson:"quantity"`
}

type Variant struct {
  Color Color `json:"color" bson:"color"`
  List []Stock `json:"list" bson:"list"`
  ImageList []string `json:"image_list" bson:"image_list"`
  ImageSqList []string `json:"image_sq_list" bson:"image_sq_list"`
}

type Product struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
  Status string `json:"status" bson:"status"`
	Code string `json:"code" bson:"code"`
  TypeId bson.ObjectId `json:"type_id" bson:"type_id,omitempty"`
  Content ContentProduct `json:"content" bson:"content"`
  Price int `json:"price" bson:"price"`
  SalePrice int `json:"sale_price" bson:"sale_price"`
	Video string `json:"video" bson:"video"`
  Image string `json:"image" bson:"image"`
	TagList []string `json:"tag_list" bson:"tag_list"`
  VariantList []Variant `json:"variant_list" bson:"variant_list"`
	LastUpdate int64 `json:"last_update" bson:"last_update"`
}

type ProductList []Product
