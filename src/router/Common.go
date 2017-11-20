package router

import (
	"net/http"
  "reflect"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"db"
	"strings"
	//"fmt"
)

func createData(w http.ResponseWriter, r *http.Request, colName string, data interface{}) {
	var text string = "Data cannot created"
	var access AccessInfo
	var err error = nil
	access, err = getAccessInfo(r)
	if err == nil {
    c := db.NewCollectionSession(access.Project.Database, colName)
		defer c.Close()

		err = c.Session.Insert(data)
		if err == nil {
			text = ""
			responseJSON(w, data, http.StatusCreated)
		}
	}

	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/createData", err)
}

func createDataWithAccess(w http.ResponseWriter, access AccessInfo, colName string, data interface{}) {
  c := db.NewCollectionSession(access.Project.Database, colName)
	defer c.Close()

	err := c.Session.Insert(data)
	if err == nil {
		responseJSON(w, data, http.StatusCreated)
	} else {
		responseWithError(w, "Data cannot created", http.StatusNotFound)
		LogServerError(access, "Router/createDataWithAccess", err)
	}
}

func createWithAccessReturn(w http.ResponseWriter, access AccessInfo, colName string, data interface{}) {
  c := db.NewCollectionSession(access.Project.Database, colName)
	defer c.Close()
	err := c.Session.Insert(data)

	if err == nil {
		responseJSON(w, data, http.StatusCreated)
	} else {
		responseWithError(w, "Data cannot created", http.StatusNotFound)
		LogServerError(access, "Router/createWithAccessReturn", err)
	}
}

func deleteData(w http.ResponseWriter, r *http.Request, colName string) {
	var text string = "Data cannot deleted"
	var id bson.ObjectId
	var access AccessInfo
	var err error = nil
	id, err = getObjectId(r)
	if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			c := db.NewCollectionSession(access.Project.Database, colName)
		 	defer c.Close()
			err = c.Session.Remove(bson.M{"_id": id})
			if err == nil {
				responseText(w, "Data has deleted", http.StatusOK)
				text = ""
			}
		}
	}

	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/deleteData", err)
}

func editData(w http.ResponseWriter, r *http.Request, colName string, data interface{}) {
	var text string = "Data cannot updated"
	var id bson.ObjectId
	var access AccessInfo
	var err error = nil
	id, err = getObjectId(r)
	if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			c := db.NewCollectionSession(access.Project.Database, colName)
			defer c.Close()
			err = c.Session.UpdateId(id, data)
			if err == nil {
				responseText(w, "Data has updated", http.StatusOK)
				text = ""
			}
		}
	}
	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/editData", err)
}

func editDataWithAccess(w http.ResponseWriter, id bson.ObjectId, access AccessInfo, colName string, data interface{}) {
	c := db.NewCollectionSession(access.Project.Database, colName)
	defer c.Close()
	err := c.Session.UpdateId(id, data)
	if err == nil {
		responseText(w, "Data has updated", http.StatusOK)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusNotFound)
		LogServerError(access, "Router/editDataWithAccess", err)
	}
}

func editWithAccessReture(w http.ResponseWriter, id bson.ObjectId, access AccessInfo, colName string, data interface{}) {
	c := db.NewCollectionSession(access.Project.Database, colName)
	defer c.Close()
	err := c.Session.UpdateId(id, data)
	if err == nil {
		responseJSON(w, data, http.StatusOK)
	} else {
		responseWithError(w, "Data cannot updated", http.StatusNotFound)
		LogServerError(access, "Router/editWithAccessReture", err)
	}
}

func getDataList(w http.ResponseWriter, r *http.Request, colName string, list interface{}, condition bson.M) {
	var text string = "Data not found"
	var access AccessInfo
	var err error = nil
	access, err = getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, colName)
		defer c.Close()

		// get list
    typ := reflect.TypeOf(list)
    nlist := reflect.New(typ).Interface()

		err = c.Session.Find(condition).All(&nlist)
		if err == nil {
			responseJsonWithError(w, nlist, http.StatusOK)
			text = ""
	 	}
	}

	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/editWithAccessReture", err)
}

func getDataItem(w http.ResponseWriter, r *http.Request, colName string, data interface{}) {
	var text string = "Data not found"
	var id bson.ObjectId
	var access AccessInfo
	var err error = nil
	id, err = getObjectId(r)
	if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			c := db.NewCollectionSession(access.Project.Database, colName)
			defer c.Close()

			typ := reflect.TypeOf(data)
	    ndata := reflect.New(typ).Interface()

			err = c.Session.Find(bson.M{"_id": id}).One(&ndata)
			if err == nil {
				responseJsonWithError(w, ndata, http.StatusOK)
				text = ""
		 	}
		}
	}

	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/editWithAccessReture", err)
}

func getItemWithConditon(w http.ResponseWriter, r *http.Request, colName string, data interface{}, condition bson.M) {
	var text string = "Data not found"
	var access AccessInfo
	var err error = nil
	access, err = getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, colName)
		defer c.Close()

		typ := reflect.TypeOf(data)
	  ndata := reflect.New(typ).Interface()

		err = c.Session.Find(condition).One(&ndata)
		if err == nil {
			responseJsonWithError(w, ndata, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusNotFound)
	LogServerError(access, "Router/editWithAccessReture", err)
}

func convertToInt(val string, def int) int {
	nval, err := strconv.Atoi(val)
	if err != nil {
		return def
	} else {
		return nval
	}
}

func replacePath(path string, access AccessInfo) string {

	return strings.Replace(path, access.Project.Address, ".", -1)
}
