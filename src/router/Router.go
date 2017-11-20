package router

import (
  "net/http"
	"github.com/gorilla/mux"
  "errors"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
)

func CreateAllRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
  addRouter(router, getProductSizeRouter())
  addRouter(router, getProductTypeRouter())
  addRouter(router, getProductRouter())
  addRouter(router, getColorRouter())
  addRouter(router, getPageRouter())
  addRouter(router, getCustomerRouter())
  addRouter(router, getUserRouter())
  addRouter(router, getOrderRouter())
  addRouter(router, getConfigurationRouter())
  addRouter(router, getSettingProjectRouter())
  addRouter(router, getLazadaRouter())
  addRouter(router, getLiveRouter())

	return router
}

func addRouter(router *mux.Router , list RouteList) {
  var temp http.HandlerFunc
  temp = handleOptions
	for _, route := range list {
    if route.RequiredToken {
      router.
  			Methods(route.Method).
  			Path(route.Pattern).
  			Name(route.Name).
  			Handler(validateToken(corsHandler(route.HandlerFunc)))
    } else {
      router.
  			Methods(route.Method).
  			Path(route.Pattern).
  			Name(route.Name).
  			Handler(corsHandler(route.HandlerFunc))
    }

    router.
  		Methods("OPTIONS").
  		Path(route.Pattern).
  		Name(route.Name + "OPTIONS").
  		Handler(corsHandler(temp))
	}
}

func getObjectId(r *http.Request) (bson.ObjectId, error) {
	objectId := bson.NewObjectId()
	vars := mux.Vars(r)
	id := vars["id"]

	if !bson.IsObjectIdHex(id) {
		err := errors.New("ID not found")
    return objectId, err
	} else {
		objectId = bson.ObjectIdHex(id)
    return objectId, nil
	}
}

func getTextId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id := vars["id"]

  return id, nil
}

func responseWithError(w http.ResponseWriter, message string, code int) {
  if message != "" {
    responseText(w, message, code)
  }
}

func responseJSON(w http.ResponseWriter, v interface{}, code int) error {
  w.WriteHeader(code)
  return json.NewEncoder(w).Encode(v)
}

func responseJsonWithError(w http.ResponseWriter, v interface{}, code int) {
  err := responseJSON(w, v, code)
  if err != nil {
    responseWithError(w, "Respond has error", code)
  }
}

func responseText(w http.ResponseWriter, message string, code int) {
  w.WriteHeader(code)
  var msg ResultMessage
  msg.Result = message
  json.NewEncoder(w).Encode(msg)
}
