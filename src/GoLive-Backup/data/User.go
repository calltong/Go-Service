package data

import (
  "sync"
  "errors"
  "golang.org/x/net/websocket"
  //"gopkg.in/mgo.v2/bson"
)

var userLock sync.Mutex

func UserRegister(user LiveUser) {
  userLock.Lock()
  val, ok := mapUser[user.Company]
  defer userLock.Unlock()
  if !ok {
    var nvar userGroup
    nvar.List = make(map[string]LiveUser)
    val = nvar
  }
  val.List[user.UserId] = user
  mapUser[user.Company] = val
}

func UserUnregister(user LiveUser) {
  userLock.Lock()
  val, ok := mapUser[user.Company]
  defer userLock.Unlock()
  if ok {
    delete(val.List, user.UserId)
    mapUser[user.Company] = val
  }
}

func NotifyAndUpdateUser(val UserNotify) error {
  userLock.Lock()
  item, ok := mapUser[val.Company]
  defer userLock.Unlock()
  if !ok {
    return errors.New("Group not found")
  }

  user, ok := item.List[val.UserId]
  if !ok {
    return errors.New("User not found")
  }

  user.AdminId = val.AdminId
  user.NetAdmin.Network = val.Network
  user.NetAdmin.IceCandidate = val.IceCandidate
  item.List[val.UserId] = user
  mapUser[val.Company] = item
  
  msg := string(val.Raw)
  if err := websocket.Message.Send(user.Connection, msg); err != nil {
    return errors.New("Can't send notify to User")
  } else {
    return nil
  }
}

func GetUser(company, userId string) (LiveUser, error) {
  var user LiveUser
  userLock.Lock()
  v, ok := mapUser[company]
  defer userLock.Unlock()
  if !ok {
    return user, errors.New("Group not found")
  }

  user, ok = v.List[userId]
  if !ok {
    return user, errors.New("User not found")
  } else {
    return user, nil
  }
}

func NotifyUser(company, userId string, data []byte) error {
  user, err := GetUser(company, userId)
  if err == nil {
    msg := string(data)
    if err := websocket.Message.Send(user.Connection, msg); err != nil {
      return errors.New("Can't send notify to User")
    } else {
      return nil
    }
  } else {
    return errors.New("User not found")
  }
}
