package database

import (
	"gopkg.in/mgo.v2/bson"
	"db"
)
var dbname string

func Initial(name string) {
	dbname = name
}

func InsertData(collection string, data interface{}) error {
	c := db.NewCollectionSession(dbname, collection)
	defer c.Close()

	return c.Session.Insert(data)
}

func UpdateData(collection string, id bson.ObjectId, data interface{}) error {
	c := db.NewCollectionSession(dbname, collection)
	defer c.Close()

  err := c.Session.UpdateId(id, data)
	return err
}

func UpsertData(collection string, id bson.ObjectId, data interface{}) error {
	c := db.NewCollectionSession(dbname, collection)
	defer c.Close()

  _, err := c.Session.UpsertId(id, data)
	return err
}

func DeleteData(collection string, id bson.ObjectId) error {
	c := db.NewCollectionSession(dbname, collection)
	return c.Session.RemoveId(id)
}
