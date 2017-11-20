package router

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	//"time"
	"db"
	"utility"
)

func getOrderRouter() RouteList {
	prefix := "/order"
	var routes = RouteList{
		Route{ "OrderCreate",   "POST", prefix + "/create", createOrder, false, },
		Route{ "OrderDelete", "DELETE", prefix + "/{id}/delete", deleteOrder, false, },
		Route{ "OrderEdit",   "PUT", prefix + "/{id}/edit", editOrder, false, },
		Route{ "OrderGetList","GET", prefix, getOrderList, false, },
		Route{ "OrderGetItem","GET", prefix + "/{id}", getOrderItem, false, },
		Route{ "OrderGetItemByCustomer","GET",  prefix + "/{id}/customer", getOrderItemByCustomer, false, },
	}

  return routes
}


func getOrderData(w http.ResponseWriter, r *http.Request) (Order, error) {
	decoder := json.NewDecoder(r.Body)
	var data Order
	err := decoder.Decode(&data)

	return data, err
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var text string = "Data cannot created"
	data, err := getOrderData(w, r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			c := db.NewCollectionSession(access.Project.Database, "Order")
			defer c.Close()
			data.Id = bson.NewObjectId()
			if data.Payment.Data.Updated == true {
				name := fmt.Sprintf("%s", data.Id.Hex())
				path := fmt.Sprintf("%s/order/%s", access.Project.Folder, data.Id.Hex())
				utility.CreatePath(path)
				temp, _ := utility.UploadImage(data.Payment.Data.Slip, path, name)
				npath := fmt.Sprintf("%s/%s", access.Project.Address, temp)
				data.Payment.Data.Slip = npath
				data.Payment.Data.Updated = false
			}
			err := c.Session.Insert(data)

			if err == nil {
				responseJSON(w, data, http.StatusCreated)
				text = ""
			}
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "Order")
}

func editOrder(w http.ResponseWriter, r *http.Request) {
	var text string = "Data cannot edited"
	data, err := getOrderData(w, r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			if data.Payment.Data.Updated == true {
				name := fmt.Sprintf("%s", data.Id.Hex())
				path := fmt.Sprintf("%s/order/%s", access.Project.Folder, data.Id.Hex())
				utility.CreatePath(path)
				temp, _ := utility.UploadImage(data.Payment.Data.Slip, path, name)
				npath := fmt.Sprintf("%s/%s", access.Project.Address, temp)
				data.Payment.Data.Slip = npath
				data.Payment.Data.Updated = false
			}

			editDataWithAccess(w, data.Id, access, "Order", data)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func getOrderList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Order")
		defer c.Close()

		queries := r.URL.Query()
		status := queries.Get("status")
		code := queries.Get("code")
		tpage := queries.Get("page")
		tlimit := queries.Get("limit")
		limit := convertToInt(tlimit, 40)
		page := convertToInt(tpage, 1) - 1
		skip := 40 * page;
		var condition bson.M
		if code != "" && status != "" {
			condition = bson.M{"status": status, "_id": bson.ObjectIdHex(code)}
		} else if code != "" {
			condition = bson.M{"_id": bson.ObjectIdHex(code)}
		} else if status != "" {
			condition = bson.M{"status": status}
		}
		// get list
		var list []Order
		err = c.Session.Find(condition).Sort("-created_at").Skip(skip).Limit(limit).All(&list)
		if err == nil {
			responseJsonWithError(w, list, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getOrderItem(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	id, err := getObjectId(r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			c := db.NewCollectionSession(access.Project.Database, "Order")
			defer c.Close()

			var data Order
			err = c.Session.Find(bson.M{"_id": id}).One(&data)
			if err == nil {
				responseJsonWithError(w, data, http.StatusOK)
				text = ""
		 	}
		}
	}

	responseWithError(w, text, http.StatusNotFound)
}

func getOrderItemByCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := getObjectId(r)
	if err == nil {
		var data Order
		getItemWithConditon(w, r, "Order", data, bson.M{"customer_id": id, "status": bson.M{"$ne": "done"}})
	} else {
		responseWithError(w, "Data not found", http.StatusInternalServerError)
	}
}
