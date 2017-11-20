package router

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"db"
)

func getColorRouter() RouteList {
	prefix := "/color"
	var routes = RouteList{
		Route{ "ColorCreate", "POST",   prefix + "/create", createColor, true, },
		Route{ "ColorDelete", "DELETE", prefix + "/{id}/delete", deleteColor, true, },
		Route{ "ColorEdit",   "PUT",  prefix + "/{id}/edit", editColor, true, },
		Route{ "ColorGetList","GET",  prefix, getColorList, false, },
		Route{ "ColorGetItem","GET",  prefix + "/{id}", getColorItem, false, },
	}

  return routes
}

type InfoColor struct {
	Name string `json:"name" bson:"name"`
}

type ContentColor struct {
	Main InfoColor `json:"main" bson:"main"`
	English InfoColor `json:"english" bson:"english"`
}

type Color struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Code string `json:"code" bson:"code"`
	Content ContentColor `json:"content" bson:"content"`
}

type ColorList []Color

func getColorData(r *http.Request) (Color, error){
	decoder := json.NewDecoder(r.Body)
	var data Color
	err := decoder.Decode(&data)

	return data, err
}

func createColor(w http.ResponseWriter, r *http.Request) {
	data, err := getColorData(r)
	if err == nil {
		createData(w, r, "Color", data)
	} else {
		responseWithError(w, "Data cannot created", http.StatusBadRequest)
	}
}

func deleteColor(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "Color")
}

func editColor(w http.ResponseWriter, r *http.Request) {
	data, err := getColorData(r)
	if err == nil {
		editData(w, r, "Color", data)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusBadRequest)
	}
}

func getColorList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Color")
	 	defer c.Close()

		// get list
		var list ColorList
		err = c.Session.Find(nil).All(&list)
		if err == nil {
			responseJsonWithError(w, list, http.StatusOK)
			text = ""
	 	}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getColorItem(w http.ResponseWriter, r *http.Request) {
	var data Color
	getDataItem(w, r, "Color", data)
}
