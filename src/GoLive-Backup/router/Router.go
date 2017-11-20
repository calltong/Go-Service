package router

import (
	"github.com/gorilla/mux"
  "common/router"
	"GoLive/router/user"
)

func CreateAllRouter() *mux.Router {
	val := mux.NewRouter().StrictSlash(true)
  router.AddRouter(val, user.GetRouter())

	return val
}
