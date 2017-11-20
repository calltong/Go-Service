package register

import (
  //"gopkg.in/mgo.v2/bson"
  "encoding/json"
  "golang.org/x/net/websocket"
  "sync"
  "errors"
  "fmt"
)

// Handle User
func HandleUser(ws *websocket.Conn) {
  fmt.Println("User connected")
  var err error
  var actor Actor
  actor.Connection = ws
  for {
    var message string
    if err = websocket.Message.Receive(ws, &message); err != nil {
      fmt.Println("User connection closed")
      break
    }

    go handleUserProcess(&actor, message)
  }
  UserUnregister(actor)
}

func handleUserProcess(user *Actor ,message string) {
  var raw []byte
  raw = getErrResponse("Message incorrect", "unknow")

  var data LiveMessage
  err := json.Unmarshal([]byte(message), &data)
  if err == nil {
    if data.Method == "request" {
      switch data.Type {
      case "register":
          raw = user.reqUserRegister(data)
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

func (actor *Actor) reqUserRegister(input LiveMessage) []byte {

  UserRegister(*actor)
  return getResponse("Register Success", input.Type)
}

var userLock sync.Mutex

func UserRegister(actor Actor) {
  userLock.Lock()
  val, ok := mapAdmin[actor.Group]
  defer userLock.Unlock()
  if !ok {
    var nvar GroupList
    nvar.List = make(map[string]Actor)
    val = nvar
  }
  val.List[actor.Id] = actor
  mapAdmin[actor.Group] = val
}

func UserUnregister(actor Actor) {
  userLock.Lock()
  val, ok := mapUser[actor.Group]
  defer userLock.Unlock()
  if ok {
    delete(val.List, actor.Id)
    mapUser[actor.Group] = val
  }
}

func UserNotifyByActor(actor Actor, raw []byte) error {
  return UserNotify(actor.Group, actor.Id, raw)
}

func UserNotify(code, id string, raw []byte) error {
  userLock.Lock()
  group, ok := mapUser[code]
  defer userLock.Unlock()
  if !ok {
    return errors.New("Group not found")
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
