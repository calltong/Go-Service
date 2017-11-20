package router

import (
	"encoding/json"
	"net/http"
	//"db"
)

func getSettingProjectRouter() RouteList {
	prefix := "/setting"
	var routes = RouteList{
		Route{ "SettingEdit",   "PUT",  prefix + "/{id}/edit", editSettingStore, true, },
		Route{ "SettingGetItem","GET",  prefix + "/{id}", getSettingStoreItem, false, },
	}

  return routes
}

func getSettingStore(w http.ResponseWriter, r *http.Request) (SettingStore, error){
	decoder := json.NewDecoder(r.Body)
	var data SettingStore
	err := decoder.Decode(&data)

	return data, err
}

func editSettingStore(w http.ResponseWriter, r *http.Request) {
	data, err := getSettingStore(w, r)
	if err == nil {
		editData(w, r, "SettingStore", data)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusBadRequest)
	}
}

func getSettingStoreItem(w http.ResponseWriter, r *http.Request) {
	var data SettingStore
	getDataItem(w, r, "SettingStore", data)
}
