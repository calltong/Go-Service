package router

import (
	"gopkg.in/mgo.v2/bson"
	"db"
	"net/http"
  //"ecommerce"
)

func getProductType(id bson.ObjectId, access AccessInfo) (ProductType, error) {
  var data ProductType
  c := db.NewCollectionSession(access.Project.Database, "ProductType")
  defer c.Close()
  err := c.Session.Find(bson.M{"_id": id}).One(&data)
  return data, err
}

func getLazadaProduct(id bson.ObjectId, access AccessInfo) (LazadaProductItem, error) {
  var data LazadaProductItem
  c := db.NewCollectionSession(access.Project.Database, "LazadaProduct")
  defer c.Close()
  err := c.Session.Find(bson.M{"_id": id}).One(&data)
  return data, err
}

func getStreetProduct(id bson.ObjectId, access AccessInfo) (StreetProductItem, error) {
  var data StreetProductItem
  c := db.NewCollectionSession(access.Project.Database, "StreetProduct")
  defer c.Close()
  err := c.Session.Find(bson.M{"_id": id}).One(&data)
  return data, err
}

func getProduct(id bson.ObjectId, access AccessInfo) (Product, error) {
  var data Product
  c := db.NewCollectionSession(access.Project.Database, "Product")
  defer c.Close()
  err := c.Session.Find(bson.M{"_id": id}).One(&data)
  return data, err
}

func getProductSize(id bson.ObjectId, access AccessInfo) (ProductSize, error) {
  var data ProductSize
	c := db.NewCollectionSession(access.Project.Database, "ProductSize")
	defer c.Close()
	err := c.Session.Find(bson.M{"_id": id}).One(&data)
	return data, err
}


func getSettingProject(access AccessInfo) (SettingStore, error) {
	c := db.NewCollectionSession(access.Project.Database, "SettingStore")
	defer c.Close()
	var data SettingStore

	err := c.Session.Find(bson.M{}).One(&data)
	return data, err
}

func getProductByRequest(r *http.Request) (bson.ObjectId, AccessInfo, Product, error) {
	var access AccessInfo
	var product Product
	id, err := getObjectId(r)
  if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			product, err = getProduct(id, access)
		}
	}
	return id, access, product, err
}
