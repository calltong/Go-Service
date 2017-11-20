package router

import (
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
	"fmt"
  "common/router"
	"GoNotify/register"
	"GoNotify/notification"
)

func CreateAllRouter() *mux.Router {
	routers := mux.NewRouter().StrictSlash(true)
	//routers.Handle("/admin", websocket.Handler(HandleTestAdmin))
  router.AddConnectionRouter(routers, register.GetRouter())
	router.AddRouter(routers, notification.GetRouter())

	return routers
}

func HandleTestAdmin(ws *websocket.Conn) {
  var err error
  for {
    var message string
    if err = websocket.Message.Receive(ws, &message); err != nil {
      fmt.Println("Admin connection closed")
      break
    }
		fmt.Println("Admin:", message)
  }

  fmt.Println("Admin have disconnection")
}
