package router

import (
  "net/http"
	"github.com/gorilla/mux"
  "errors"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "common/authentication"
)

type TotalMessage struct {
	Value int `json:"value"`
}

type ResultMessage struct {
	Result string `json:"result"`
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Func http.HandlerFunc
  RequiredToken bool
}

type RouteList []Route

type RouteConnection struct {
	Name        string
	Pattern     string
	Func http.Handler
}

type RouteConnectionList []RouteConnection

func AddRouter(router *mux.Router , list RouteList) {
  var temp http.HandlerFunc
  temp = handleOptions
	for _, route := range list {
    if route.RequiredToken {
      router.
  			Methods(route.Method).
  			Path(route.Pattern).
  			Name(route.Name).
  			Handler(authentication.ValidateToken(CorsHandler(route.Func)))
    } else {
      router.
  			Methods(route.Method).
  			Path(route.Pattern).
  			Name(route.Name).
  			Handler(CorsHandler(route.Func))
    }

    router.
  		Methods("OPTIONS").
  		Path(route.Pattern).
  		Name(route.Name + "OPTIONS").
  		Handler(CorsHandler(temp))
	}
}

func AddConnectionRouter(router *mux.Router , list RouteConnectionList) {
	for _, route := range list {
    router.
      Path(route.Pattern).
      Name(route.Name).
      Handler(route.Func)
	}
}

func validateToken(page http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    page(res, req)
	})
}

func GetObjectId(r *http.Request) (bson.ObjectId, error) {
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

func GetTextId(r *http.Request) (string, error) {
	vars := mux.Vars(r)
	id := vars["id"]

  return id, nil
}

func ResponseWithError(w http.ResponseWriter, message string, code int) {
  if message != "" {
    ResponseText(w, message, code)
  }
}

func ResponseJSON(w http.ResponseWriter, v interface{}, code int) error {
  w.WriteHeader(code)
  return json.NewEncoder(w).Encode(v)
}

func ResponseJsonWithError(w http.ResponseWriter, v interface{}, code int) {
  err := ResponseJSON(w, v, code)
  if err != nil {
    ResponseWithError(w, "Respond has error", code)
  }
}

func ResponseText(w http.ResponseWriter, message string, code int) {
  w.WriteHeader(code)
  var msg ResultMessage
  msg.Result = message
  json.NewEncoder(w).Encode(msg)
}

func ResponseSuccess(w http.ResponseWriter) {
  w.WriteHeader(http.StatusOK)
  var msg ResultMessage
  msg.Result = "Success"
  json.NewEncoder(w).Encode(msg)
}

func ResponseJsonSuccess(w http.ResponseWriter, v interface{}) error {
  w.WriteHeader(http.StatusOK)
  return json.NewEncoder(w).Encode(v)
}
