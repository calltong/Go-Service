package user

import (
	"encoding/json"
	"net/http"
	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	//"fmt"
	"common/router"
	"GoLive/common/database"
)

var colUser string = "user"

func GetRouter() router.RouteList {
	var routes = router.RouteList {
		//router.Route{ "UserOnline", "PUT", "/user/status/{company}/{id}/online", statusOnline, false, },
		router.Route{ "UserGenerateId", "GET", "/user/id/generate", generateUserId, false, },
		router.Route{ "UserOnline", "POST", "/user/connect", connectToAdmin, false, },
		router.Route{ "UserBrowser", "POST", "/user/{id}/browser", updateBrowser, false, },
	}

  return routes
}

func getConnectMessage(r *http.Request) (ReqConnect, error){
	decoder := json.NewDecoder(r.Body)
	var data ReqConnect
	err := decoder.Decode(&data)

	return data, err
}

func getUserData(r *http.Request) (UserData, error){
	decoder := json.NewDecoder(r.Body)
	var data UserData
	err := decoder.Decode(&data)

	return data, err
}

func generateUserId(w http.ResponseWriter, r *http.Request) {
	var data ResUserId
	data.UserId = bson.NewObjectId()

	router.ResponseJSON(w, data, http.StatusOK)
}

func updateBrowser(w http.ResponseWriter, r *http.Request) {
	data, err := getUserData(r)

	if err != nil {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := router.GetObjectId(r)
	if err != nil {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
	}

  upsertData := bson.M{ "$set": bson.M {"browser": data.Browser}}
	err = database.UpsertData(colUser, id, upsertData)
	if err == nil {
		router.ResponseSuccess(w)
	} else {
		router.ResponseText(w, err.Error(), http.StatusBadRequest)
	}
}

func getReadyAdmin(company string) (string, error){
	return "", nil
}

func connectToAdmin(w http.ResponseWriter, r *http.Request) {
	input, err := getConnectMessage(r)
	if err != nil {
    router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
  }

	_, err = getReadyAdmin(input.Company)
  if err != nil {
    router.ResponseText(w, err.Error(), http.StatusBadRequest)
		return
  }

	router.ResponseWithError(w, "Success", http.StatusOK)
}
