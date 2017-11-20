package router

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"time"
	"db"
)

func getCustomerRouter() RouteList {
	prefix := "/customer"
	var routes = RouteList {
		Route{ "CustomerCreate", "POST",   prefix + "/create", createCustomer, false, },
		Route{ "CustomerDelete", "DELETE", prefix + "/{id}/delete", deleteCustomer, false, },
		Route{ "CustomerEdit",   "PUT",  prefix + "/{id}/edit", editCustomer, false, },
		Route{ "CustomerGetList","GET",  prefix, getCustomerList, false, },
		Route{ "CustomerGetItem","GET",  prefix + "/{id}/id", getCustomerItem, false, },
		Route{ "CustomerGetItem","GET",  prefix + "/{id}/browser", getCustomerByBrowser, false, },
	}

  return routes
}

func getCustomerData(w http.ResponseWriter, r *http.Request) (Customer, error) {
	decoder := json.NewDecoder(r.Body)
	var data Customer
	err := decoder.Decode(&data)

	return data, err
}

func getCustomerByBrowser(w http.ResponseWriter, r *http.Request) {
	var text string = "Id not found"
	id, err := getTextId(r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			var data Customer
			c := db.NewCollectionSession(access.Project.Database, "Customer")
			defer c.Close()

			err = c.Session.Find(bson.M{"browser": id}).One(&data)
			if err == nil {
				responseJsonWithError(w, data, http.StatusOK)
				text = ""
			}
		}
	}

	if text != "" {
		responseText(w, text, http.StatusNotFound)
	}
}

func getCustomerByEmail(access AccessInfo, email string, typ string) (Customer, error) {
	var condition bson.M
	condition = bson.M{"email": email, "type": typ}
	var data Customer
	c := db.NewCollectionSession(access.Project.Database, "Customer")
	defer c.Close()

	err := c.Session.Find(condition).One(&data)
	return data, err
}

func getCustomerByMedia(access AccessInfo, customer Customer) (Customer, error) {
	var data Customer
	var err error
	c := db.NewCollectionSession(access.Project.Database, "Customer")
	defer c.Close()
	for _, item := range customer.MediaList {
		err = c.Session.Find(bson.M{"media_list.type": item.Type, "media_list.id": item.Id}).One(&data)
		if err == nil {
			return data, nil
		}
	}

	return data, err
}

func GetCustomer(access AccessInfo, customer Customer) (Customer, error) {
	if customer.Email != "" {
		data, err := getCustomerByEmail(access, customer.Email, customer.Type)
		if err == nil {
			return data, err
		}
	}

	data, err := getCustomerByMedia(access, customer)
	return data, err
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	text := "Data cannot created"
	data, err := getCustomerData(w, r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			_, err := GetCustomer(access, data)

			if err == nil {
				data.Id = bson.NewObjectId()
				data.LastUpdated = time.Now().Unix()
				createWithAccessReturn(w, access, "Customer", data)
				text = ""
			} else {
				text = "Account already register"
			}
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "Customer")
}

func editCustomer(w http.ResponseWriter, r *http.Request) {
	data, err := getCustomerData(w, r)
	if err == nil {
		editData(w, r, "Customer", data)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusBadRequest)
	}
}

func getCustomerList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Customer")
		defer c.Close()

		// get list
		var list CustomerList
		err = c.Session.Find(nil).All(&list)
		if err == nil {
			responseJsonWithError(w, list, http.StatusOK)
			text = ""
	 	}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getCustomerItem(w http.ResponseWriter, r *http.Request) {
	var data Customer
	getDataItem(w, r, "Customer", data)
}
