package admin

import (
	//"encoding/json"
	"net/http"
	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	//"fmt"
	"common/router"
	"GoLive/common"
	"GoLive/common/database"
	"config"
	"db"
)

var colAdmin string = "admin"
var colActiveAdmin string = "active_admin"

func GetRouter() router.RouteList {
	var routes = router.RouteList {
		router.Route{ "AdminLogin", "PUT", "/admin/login", loginAdmin, false, },
		router.Route{ "AdminOnline", "PUT", "/admin/status/{id}/online", setOnline, false, },
		router.Route{ "AdminOffline", "PUT", "/admin/status/{id}/offline", setOnline, false, },
	}

  return routes
}

func loginAdmin(w http.ResponseWriter, r *http.Request) {
	login, err := common.GetLoginData(r)
	if err != nil {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbname := config.GetDatabaseName()
	var data AdminData
  c := db.NewCollectionSession(dbname, colAdmin)
  defer c.Close()
  err = c.Session.Find(bson.M{"username": login.Username, "password": login.Password}).One(&data)
  if err == nil {
		router.ResponseJsonSuccess(w, data)
	} else {
		router.ResponseText(w, err.Error(), http.StatusNotFound)
	}
}

func setOnline(w http.ResponseWriter, r *http.Request) {
	id, err := router.GetObjectId(r)
	if err != nil {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbname := config.GetDatabaseName()
	var data AdminData
  c := db.NewCollectionSession(dbname, colAdmin)
  defer c.Close()
  err = c.Session.Find(bson.M{"_id": id}).One(&data)
  if err != nil {
		router.ResponseText(w, err.Error(), http.StatusNotFound)
	}

	var admin ActiveAdminData
	admin.Id = data.Id
	admin.GroupCode = data.GroupCode
	admin.Company = data.Company
	admin.Name = data.Name
	admin.Image = data.Image
	admin.Maximun = data.Maximun
	admin.Counter = 0

	err = database.UpsertData(colActiveAdmin, id, admin)

	if err == nil {
		router.ResponseSuccess(w)
	} else {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
	}
}

func setOffline(w http.ResponseWriter, r *http.Request) {
	id, err := router.GetObjectId(r)
	if err != nil {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DeleteData(colActiveAdmin, id)
	if err == nil {
		router.ResponseSuccess(w)
	} else {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
	}
}
