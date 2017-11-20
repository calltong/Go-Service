package router

import (
  "golang.org/x/net/websocket"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
	"fmt"
  "GoLive/data"
)

type GoLiveUser struct {
  Connection *websocket.Conn
  SessionId string
  Company string
  UserId string
}

func HandleUser(ws *websocket.Conn) {
  var err error
  var user GoLiveUser
  user.Connection = ws
  for {
    var message string
    if err = websocket.Message.Receive(ws, &message); err != nil {
      fmt.Println("User connection closed")
      break
    }

    go handleUserProcess(&user, message)
  }
}

func handleUserProcess(user *GoLiveUser ,message string) {
  var raw []byte
  raw = getErrResponse("Message incorrect", "unknow")

  var data LiveMessage
  err := json.Unmarshal([]byte(message), &data)
  if err == nil {
    if data.Method == "request" {
      switch data.Type {
        case "connection":
          raw = user.userReqConnection(data)
          break;
        case "reconnection":
          raw = user.userReqReconnection(data)
          break;
      }
    }
  }
  if len(raw) > 0 {
    reply := string(raw)
    if err = websocket.Message.Send(user.Connection, reply); err != nil {
      fmt.Println("User can't send")
    }
  }
}

func (user *GoLiveUser) userReqConnection(input LiveMessage) []byte {
  typ := input.Type
  var connect ConnectReq
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &connect)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }
  fmt.Println("User Connection", user.UserId)
  SessionId := bson.NewObjectId().Hex()
  user.UserId = connect.UserId
  user.Company = connect.Company

  var live data.LiveUser
  live.UserId = connect.UserId
  live.Company = connect.Company
  live.Connection = user.Connection
  live.NetUser.Network = connect.NetUser.Network
  live.NetUser.IceCandidate = connect.NetUser.IceCandidate

  data.UserRegister(live)

  admin, err := data.ConnectAdmin(connect.Company)
  if err != nil {
    return getErrResponse("Admin is not ready", typ)
  }

  err = notifyNewUserToAdmin(admin, connect, SessionId)
  if err == nil {
    var nores []byte
    return nores
  } else {
    return getErrResponse("Admin notify new user fail", typ)
  }
}

func (user *GoLiveUser) userReqReconnection(input LiveMessage) []byte {
  typ := input.Type
  var connect ReconnectReq
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &connect)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }
  fmt.Println("User Connection", connect.UserId)
  _, err = data.GetAdmin(connect.Company, connect.AdminId)
  if err != nil {
    _, err = data.ConnectAdmin(connect.Company)
    if err != nil {
      return getErrResponse("Admin is not ready", typ)
    }
  }

  _, err = data.GetUser(connect.Company, connect.UserId)


  var message LiveMessage
  message.Method = "request"
  message.Type = "newuser"
  //message.Data = bs
  raw, _ = json.Marshal(message)
  return raw
}

func notifyNewUserToAdmin(admin data.LiveAdmin, connect ConnectReq, sessionId string) error {
  var nUser UserNotifyReq
  nUser.SessionId = sessionId
  nUser.Company = connect.Company
  nUser.UserId = connect.UserId
  nUser.NetUser = connect.NetUser
  //nUser.IceCandidate = connect.NetUser.IceCandidate

  var bs bson.M
  raw, _ := bson.Marshal(nUser)
  bson.Unmarshal(raw, &bs)

  raw = buildMessageReq("newuser", bs)

  return data.NotifyAdmin(admin.Company, admin.AdminId, raw)
}
