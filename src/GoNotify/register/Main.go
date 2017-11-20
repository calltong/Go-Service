package register

import (
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "golang.org/x/net/websocket"
  "common/router"
)

var mapAdmin map[string]GroupList
var mapUser map[string]GroupList

func Init() {
  mapAdmin = make(map[string]GroupList)
  mapUser = make(map[string]GroupList)
}

func GetRouter() router.RouteConnectionList {
	var routes = router.RouteConnectionList {
		router.RouteConnection{ "RegisterUser", "/user", websocket.Handler(HandleUser), },
    router.RouteConnection{ "RegisterAdmin", "/admin", websocket.Handler(HandleAdmin), },
	}

  return routes
}

func getErrResponse(input, typ string) []byte {
  var res ResponseMessage
  res.Method = "response"
  res.Type = typ
  res.Message = input
  res.Status = 400;
  raw, _ := json.Marshal(res)
  return raw
}

func getResponse(input, typ string) []byte {
  var res ResponseMessage
  res.Method = "response"
  res.Type = typ
  res.Message = input
  res.Status = 200;
  raw, _ := json.Marshal(res)
  return raw
}

func buildMessageReq(typ string, bs bson.M) []byte {
  var res LiveMessage
  res.Method = "request"
  res.Type = typ
  res.Data = bs
  res.Status = 200;
  raw, _ := json.Marshal(res)
  return raw
}

func buildMessageRes(typ string, bs bson.M) []byte {
  var res LiveMessage
  res.Method = "response"
  res.Type = typ
  res.Data = bs
  res.Status = 200;
  raw, _ := json.Marshal(res)
  return raw
}
