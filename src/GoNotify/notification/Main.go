package notification

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	//"gopkg.in/mgo.v2/bson"
	"fmt"
	"common/router"
	"GoNotify/register"
)

func GetRouter() router.RouteList {
	var routes = router.RouteList {
		router.Route{ "NotifyUser", "PUT", "/user/notify/{company}/{id}", notifyToUser, false, },
		router.Route{ "NotifyAdmin", "PUT", "/admin/notify/{company}/{id}", notifyToAdmin, false, },
	}

  return routes
}

func getMessage(r *http.Request) (NotifyMessage, error){
	decoder := json.NewDecoder(r.Body)
	var data NotifyMessage
	err := decoder.Decode(&data)

	return data, err
}

func getParameter(r *http.Request) (string, string) {
	vars := mux.Vars(r)
	company := vars["company"]
	id := vars["id"]
	return company, id
}

func notifyToUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("notify to user")
	input, err := getMessage(r)
	if err != nil {
		router.ResponseWithError(w, "Body incorrect", http.StatusBadRequest)
		return
	}

	raw, _ := json.Marshal(input.Message)
	company, id := getParameter(r)
	register.UserNotify(company, id, raw)
	router.ResponseSuccess(w)
}

func notifyToAdmin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("notify to admin")
	input, err := getMessage(r)
	if err != nil {
		router.ResponseWithError(w, "Body incorrect", http.StatusBadRequest)
		return
	}
	raw, _ := json.Marshal(input.Message)
	company, id := getParameter(r)

	register.AdminNotify(company, id, raw)
	router.ResponseSuccess(w)
}
