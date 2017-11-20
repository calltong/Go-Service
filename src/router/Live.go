package router

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"db"
)

func getLiveRouter() RouteList {
	prefix := "/live"
	var routes = RouteList{
		Route{ "LiveUserRegister", "PUT", prefix + "/user/{id}/register", liveUserRegister, false, },
		Route{ "LiveUserUnregister", "PUT", prefix + "/user/{id}/unregister", liveUserUnregister, false, },

		Route{ "LiveUserGetAdmin", "GET", prefix + "/user/get/online", liveGetAdminReady, false, },

    Route{ "LiveAdminRegister", "PUT", prefix + "/admin/register", liveAdminRegister, false, },
    Route{ "LiveAdminUnregister", "PUT", prefix + "/admin/{id}/unregister", liveAdminUnregister, false, },
		Route{ "LiveGetById", "GET", prefix + "/{id}/get", LiveGetById, false, },
	}

  return routes
}

type LiveResAdmin struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Owner LiveNetwork `json:"owner" bson:"owner"`
}

type LiveNetwork struct {
	Status string `json:"status" bson:"status"`
	Network bson.M `json:"network" bson:"network"`
	IceCandidate bson.M `json:"candidate" bson:"candidate"`
}

type LiveData struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Name string `json:"name" bson:"name"`
	Status string `json:"status" bson:"status"`
  Type string `json:"type" bson:"type"`
	UID string `json:"uid" bson:"uid"`
	Owner LiveNetwork `json:"owner" bson:"owner"`
	Client LiveNetwork `json:"client" bson:"client"`
}

func LiveGetById(w http.ResponseWriter, r *http.Request) {
	id, err := getObjectId(r)
	if err == nil {
		c := db.NewCollectionSession("pompom_db", "LiveUser")
		defer c.Close()

		var data LiveData
		err := c.Session.Find(bson.M{"_id": id}).One(&data)
		if err == nil {
			responseJsonWithError(w, data, http.StatusOK)
		} else {
			responseText(w, "No User", http.StatusNotFound)
		}
	} else {
		responseText(w, "No User Id", http.StatusNotFound)
	}
}

func liveGetAdminReady(w http.ResponseWriter, r *http.Request) {
	c := db.NewCollectionSession("pompom_db", "LiveUser")
	defer c.Close()

	var data LiveData
	err := c.Session.Find(bson.M{"status": "online"}).One(&data)
	if err == nil {
		var live LiveResAdmin
		live.Id = data.Id
		live.Name = data.Name
		live.Owner = data.Owner
		responseJsonWithError(w, live, http.StatusOK)
	} else {
		responseText(w, "No Admin User", http.StatusNotFound)
	}
}

func liveUserRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var network LiveNetwork
	err := decoder.Decode(&network)
	if err == nil {
		id, err := getObjectId(r)

		if err == nil {
			c := db.NewCollectionSession("pompom_db", "LiveUser")
			defer c.Close()

			var data LiveData
			err = c.Session.Find(bson.M{"_id": id}).One(&data)

			data.Client.Network = network.Network
			data.Client.IceCandidate = network.IceCandidate
			data.Client.Status = "ready"
			err = c.Session.UpdateId(data.Id, data)
			responseJsonWithError(w, data, http.StatusOK)
		} else {
			responseText(w, "No User Id", http.StatusNotFound)
		}
	} else {
		responseText(w, "No Description", http.StatusNotFound)
	}
}

func liveUserUnregister(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, "Data cannot created", http.StatusBadRequest)
}

func liveAdminRegister(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var network LiveNetwork
	err := decoder.Decode(&network)
	if err == nil {
		c := db.NewCollectionSession("pompom_db", "LiveUser")
		defer c.Close()

		var data LiveData
		err = c.Session.Find(bson.M{"username": "calltong" , "password": "324661"}).One(&data)

		data.Status = "online"
		data.Owner.Network = network.Network
		data.Owner.IceCandidate = network.IceCandidate
		data.Owner.Status = "ready"
		err = c.Session.UpdateId(data.Id, data)
		responseJsonWithError(w, data, http.StatusOK)
	} else {
		responseText(w, "No Description", http.StatusNotFound)
	}
}

func liveAdminUnregister(w http.ResponseWriter, r *http.Request) {
	id, err := getObjectId(r)
	if err == nil {
		c := db.NewCollectionSession("pompom_db", "LiveUser")
		defer c.Close()

		var data LiveData
		err := c.Session.Find(bson.M{"_id": id}).One(&data)
	  if err == nil {
			data.Status = "offline"
			err = c.Session.UpdateId(data.Id, data)
			responseText(w, "Data has completed", http.StatusOK)
		} else {
			responseText(w, "No Description", http.StatusNotFound)
		}
	} else {
		responseText(w, "No User Id", http.StatusNotFound)
	}

}
