package router

import (
	"db"
)

func insertLazadaProduct(access AccessInfo, product LazadaProductItem) error {
  c := db.NewCollectionSession(access.Project.Database, "LazadaProduct")
	defer c.Close()

	err := c.Session.Insert(product)
  return err
}

func updateLazadaProduct(access AccessInfo, product LazadaProductItem) error {
  c := db.NewCollectionSession(access.Project.Database, "LazadaProduct")
	defer c.Close()

	err := c.Session.UpdateId(product.Id, product)
  return err
}

func insertStreetProduct(access AccessInfo, product StreetProductItem) error {
  c := db.NewCollectionSession(access.Project.Database, "StreetProduct")
	defer c.Close()

	err := c.Session.Insert(product)
  return err
}

func updateStreetProduct(access AccessInfo, product StreetProductItem) error {
  c := db.NewCollectionSession(access.Project.Database, "StreetProduct")
	defer c.Close()

	err := c.Session.UpdateId(product.Id, product)
  return err
}
