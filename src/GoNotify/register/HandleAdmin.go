package register

import (
  "gopkg.in/mgo.v2/bson"
  "errors"
  "encoding/json"
  "golang.org/x/net/websocket"
  "sync"
  "fmt"
)

// Handle Admin
func HandleAdmin(ws *websocket.Conn) {
  fmt.Println("Admin connected")
  var err error
  var actor Actor
  actor.Connection = ws
  for {
    var message string
    if err = websocket.Message.Receive(ws, &message); err != nil {
      fmt.Println("Admin connection closed")
      break
    }

    go handleAdminProcess(&actor, message)
  }
  AdminUnregister(actor)
}

func handleAdminProcess(user *Actor ,message string) {
  var raw []byte
  raw = getErrResponse("Message incorrect", "unknow")

  var data LiveMessage
  err := json.Unmarshal([]byte(message), &data)
  if err == nil {
    if data.Method == "request" {
      switch data.Type {
      case "register":
          raw = user.reqAdminRegister(data)
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

func (actor *Actor) reqAdminRegister(input LiveMessage) []byte {
  typ := input.Type
  var msg ReqRegister
  raw, _ := bson.Marshal(input.Data)
	err := bson.Unmarshal(raw, &msg)
  if err != nil {
    return getErrResponse("Data package incorrect", typ)
  }

  actor.Id = msg.Id
  actor.Group = msg.Group

  AdminRegister(*actor)
  return getResponse("Register Success", input.Type)
}

var adminLock sync.Mutex

func AdminRegister(actor Actor) {
  adminLock.Lock()
  val, ok := mapAdmin[actor.Group]
  defer adminLock.Unlock()
  if !ok {
    var nvar GroupList
    nvar.List = make(map[string]Actor)
    val = nvar
  }
  val.List[actor.Id] = actor
  mapAdmin[actor.Group] = val
}

func AdminUnregister(actor Actor) {
  adminLock.Lock()
  val, ok := mapAdmin[actor.Group]
  defer adminLock.Unlock()
  if ok {
    delete(val.List, actor.Id)
    mapAdmin[actor.Group] = val
  }
}

func AdminNotifyByActor(actor Actor, raw []byte) error {
  return AdminNotify(actor.Group, actor.Id, raw)
}

func AdminNotify(gp, id string, raw []byte) error {
  adminLock.Lock()
  group, ok := mapAdmin[gp]
  defer adminLock.Unlock()
  if !ok {
    return errors.New("Company not found")
  }

  user, ok := group.List[id]
  if !ok {
    return errors.New("User not found")
  }

  if len(raw) > 0 {
    msg := string(raw)
    if err := websocket.Message.Send(user.Connection, msg); err != nil {
      return errors.New("Can not send message")
    }
  }

  return nil
}
