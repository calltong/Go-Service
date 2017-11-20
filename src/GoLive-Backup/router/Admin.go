package router

import (
  "golang.org/x/net/websocket"
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
	"fmt"
  "db"
  "GoLive/data"
)

type GoLiveAdmin struct {
  AdminId string
  Connection *websocket.Conn
}

func HandleAdmin(ws *websocket.Conn) {
  var err error
  var admin GoLiveAdmin
  admin.Connection = ws
  admin.AdminId = "ini"
  for {
    var message string
    if err = websocket.Message.Receive(ws, &message); err != nil {
      fmt.Println("Admin connection closed")
      break
    }

    go handleAdminProcess(&admin, message)
  }

  fmt.Println("Admin have disconnection")
}

func handleAdminProcess(admin *GoLiveAdmin, message string) {
  var raw []byte
  raw = getErrResponse("Message incorrect", "unknow")

  var data LiveMessage
  err := json.Unmarshal([]byte(message), &data)
  if err == nil {
    if data.Method == "request" {
      switch data.Type {
        case "register":
          fmt.Println("Admin Req Register")
          raw = admin.adminReqRegister(data)
          //fmt.Println("Admin have register:", id)
          break;
      }
    } else if data.Method == "response" {
      switch data.Type {
        case "newuser":
          fmt.Println("Admin Res New User")
          raw = admin.adminResUserNotify(data)
          break;
      case "test":
        raw = admin.adminResTest(data)
        break;
      }
    }
  }

  if len(raw) > 0 {
    reply := string(raw)
    if err = websocket.Message.Send(admin.Connection, reply); err != nil {
      fmt.Println("Admin can't send")
    }
  }
}

func (admin *GoLiveAdmin) adminReqRegister(input LiveMessage) []byte {
  typ := input.Type
  var msg AdminRegisterReq
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &msg)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }

  if !bson.IsObjectIdHex(msg.AdminId) {
    return getErrResponse("Id incorrect", typ)
  }

  c := db.NewCollectionSession("golive_db", "admin")
  defer c.Close()

  var val AdminData
  id := bson.ObjectIdHex(msg.AdminId)
  err = c.Session.FindId(id).One(&val)
  if err != nil {
    return getErrResponse("Id not found", typ)
  }

  var live data.LiveAdmin
  live.AdminId = msg.AdminId
  live.Company = val.Company
  live.Name = val.Name
  live.Connection = admin.Connection
  live.Max = 5
  live.Count = 0
  data.AdminRegister(live)

  admin.AdminId = msg.AdminId
  fmt.Println("Admin id:", admin.AdminId)
  return getResponse("Success", typ)
}

func (admin *GoLiveAdmin) unregisterAdmin(input LiveMessage) []byte {
  typ := input.Type
  var val AdminRegisterReq
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &val)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }

  //unregisterAdminById(data.Id)
  return getResponse("Success", typ)
}

func (admin *GoLiveAdmin) unregisterAdminById(company, id string) {

}

func (admin *GoLiveAdmin) adminResUserNotify(input LiveMessage) []byte {
  typ := input.Type
  var nUser UserNotifyRes
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &nUser)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }

  var res ConnectRes
  res.SessionId = nUser.SessionId
  res.AdminId = admin.AdminId
  res.NetAdmin = nUser.NetAdmin

  var bs bson.M
  raw, _ = bson.Marshal(res)
  bson.Unmarshal(raw, &bs)

  raw = buildMessageRes("connection", bs)
  var notify data.UserNotify
  notify.Company = nUser.Company
  notify.UserId = nUser.UserId
  notify.AdminId = admin.AdminId
  notify.Network = nUser.NetAdmin.Network
  notify.IceCandidate = nUser.NetAdmin.IceCandidate
  notify.Raw = raw

  fmt.Println("Res Admin id:", admin.AdminId)
  err = data.NotifyAndUpdateUser(notify)

  if err == nil {
    var nores []byte
    return nores
  } else {
    return getErrResponse("User notify acception fail", typ)
  }
}

func (admin *GoLiveAdmin) adminResTest(input LiveMessage) []byte {
  typ := input.Type
  var nUser UserNotifyRes
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &nUser)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }
  fmt.Println("user: ", nUser.UserId)
  user, err := data.GetUser("COCY2017", nUser.UserId)
  user.Done <- true
  var nores []byte
  return nores
}
