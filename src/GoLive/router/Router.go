package router

import (
	"github.com/gorilla/mux"
  "common/router"
	"GoLive/router/user"
	"GoLive/router/admin"
)

func CreateAllRouter() *mux.Router {
	routers := mux.NewRouter().StrictSlash(true)
  router.AddRouter(routers, user.GetRouter())
	router.AddRouter(routers, admin.GetRouter())

	return routers
}
