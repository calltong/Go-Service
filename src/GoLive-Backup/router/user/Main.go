package user

import (
	"net/http"
	//"encoding/json"
  "gopkg.in/mgo.v2/bson"
	"fmt"
	"common/router"
  "GoLive/data"
	"GoLive/notify"
)

func GetRouter() router.RouteList {
	prefix := "/user"
	var routes = router.RouteList {
		router.Route{ "UserConnection",  "POST",  prefix + "/connection", getUserConnection, false, },
	}

  return routes
}

func getUserConnection(w http.ResponseWriter, r *http.Request) {
	var text string = "connection"

	var live data.LiveUser
  live.UserId = "A12345"
  live.Company = "COCY2017"
  data.UserRegister(live)

  admin, err := data.ConnectAdmin("COCY2017")
	var nUser TestData
	nUser.Method = "request"
	nUser.Type = "test"
	nUser.UserId = "A12345"

  raw, _ := bson.Marshal(nUser)
	fmt.Println("admin:", admin)
  err = notify.NotifyToAdmin(live, admin, raw)
	fmt.Println("err:", err)
	router.ResponseWithError(w, text, http.StatusOK)
}
