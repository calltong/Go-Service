package router

import (
	"encoding/json"
	//"fmt"
	"net/http"
	"db"
)

func getProductTypeRouter() RouteList {
	prefix := "/protype"
	var routes = RouteList{
		Route{ "TypeCreate", "POST",   prefix + "/create", createProductType, true, },
		Route{ "TypeDelete", "DELETE", prefix + "/{id}/delete", deleteProductType, true, },
		Route{ "TypeEdit",   "PUT",  prefix + "/{id}/edit", editProductType, true, },
		Route{ "TypeGetList","GET",  prefix, getProductTypeList, false, },
		Route{ "TypeGetItem","GET",  prefix + "/{id}", getProductTypeItem, false, },
	}

  return routes
}

func getProductTypeData(w http.ResponseWriter, r *http.Request) (ProductType, error) {
	decoder := json.NewDecoder(r.Body)
	var data ProductType
	err := decoder.Decode(&data)

	return data, err
}

func createProductType(w http.ResponseWriter, r *http.Request) {
	data, err := getProductTypeData(w, r)
	if err == nil {
		createData(w, r, "ProductType", data)
	} else {
		responseWithError(w, "Data cannot created", http.StatusBadRequest)
	}
}

func deleteProductType(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "ProductType")
}

func editProductType(w http.ResponseWriter, r *http.Request) {
	data, err := getProductTypeData(w, r)
	if err == nil {
		editData(w, r, "ProductType", data)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusBadRequest)
	}
}

func getProductTypeList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "ProductType")
		defer c.Close()
		// get list
		var list ProductTypeList
		err = c.Session.Find(nil).All(&list)
		if err == nil {
			responseJsonWithError(w, list, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getProductTypeItem(w http.ResponseWriter, r *http.Request) {
	var data ProductType
	getDataItem(w, r, "ProductType", data)
}
