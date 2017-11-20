package router

import (
	"encoding/json"
	"net/http"
	"db"
)

func getProductSizeRouter() RouteList {
	prefix := "/prosize"
	var routes = RouteList{
		Route{ "SizeCreate", "POST",   prefix + "/create", createProductSize, true, },
		Route{ "SizeDelete", "DELETE", prefix + "/{id}/delete", deleteProductSize, true, },
		Route{ "SizeEdit",   "PUT",  prefix + "/{id}/edit", editProductSize, true, },
		Route{ "SizeGetList","GET",  prefix, getProductSizeList, false, },
		Route{ "SizeGetItem","GET",  prefix + "/{id}", getProductSizeItem, false, },
	}

  return routes
}

func getProductSizeData(w http.ResponseWriter, r *http.Request) (ProductSize, error){
	decoder := json.NewDecoder(r.Body)
	var data ProductSize
	err := decoder.Decode(&data)

	return data, err
}

func createProductSize(w http.ResponseWriter, r *http.Request) {
	data, err := getProductSizeData(w, r)
	if err == nil {
		createData(w, r, "ProductSize", data)
	} else {
		responseWithError(w, "Data cannot created", http.StatusBadRequest)
	}
}

func deleteProductSize(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "ProductSize")
}

func editProductSize(w http.ResponseWriter, r *http.Request) {
	data, err := getProductSizeData(w, r)
	if err == nil {
		editData(w, r, "ProductSize", data)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusBadRequest)
	}
}

func getProductSizeList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "ProductSize")
		defer c.Close()
		// get list
		var list ProductSizeList
		err = c.Session.Find(nil).All(&list)
		if err == nil {
			responseJsonWithError(w, list, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getProductSizeItem(w http.ResponseWriter, r *http.Request) {
	var data ProductSize
	getDataItem(w, r, "ProductSize", data)
}
