package data

import (
  "sync"
  "errors"
  "fmt"
  "golang.org/x/net/websocket"
)

var adminLock sync.Mutex

func AdminRegister(admin LiveAdmin) {
  adminLock.Lock()
  val, ok := mapAdmin[admin.Company]
  defer adminLock.Unlock()
  if !ok {
    var nvar adminGroup
    nvar.List = make(map[string]LiveAdmin)
    val = nvar
  }
  val.List[admin.AdminId] = admin
  mapAdmin[admin.Company] = val
}

func AdminUnregister(admin LiveAdmin) {
  adminLock.Lock()
  val, ok := mapAdmin[admin.Company]
  defer adminLock.Unlock()
  if ok {
    delete(val.List, admin.AdminId)
    mapAdmin[admin.Company] = val
  }
}

func ConnectAdmin(company string) (LiveAdmin, error) {
  var err error
  var admin LiveAdmin
  adminLock.Lock()
  v, ok := mapAdmin[company]
  fmt.Println("Connect:", company, ok)
  defer adminLock.Unlock()
  if ok {
    for key := range v.List {
      admin := v.List[key]
      fmt.Println("Connect:", admin.Max, admin.Count)
      if admin.Max != admin.Count {
        admin.Count += 1
        v.List[key] = admin
        mapAdmin[company] = v
        return admin, nil
      }
    }
    err = errors.New("Admin not found")
  } else {
    err = errors.New("Group not found")
  }

  return admin, err
}

func DisconnectAdmin(company string, id string) error {
  var err error
  adminLock.Lock()
  v, ok := mapAdmin[company]
  defer adminLock.Unlock()
  if ok {
    err = errors.New("Group not found")
  } else {
    val, okval := v.List[id]
    if okval {
      val.Count -= 1
      if val.Count < 0 {
        val.Count = 0
      }
      v.List[id] = val
      mapAdmin[company] = v
    }
  }

  return err
}

func GetAdmin(company, id string) (LiveAdmin, error) {
  var err error
  var admin LiveAdmin
  adminLock.Lock()
  v, ok := mapAdmin[company]
  defer adminLock.Unlock()
  if ok {
    admin, ok = v.List[id]
    if ok {
      return admin, nil
    }
    err = errors.New("Admin not found")
  } else {
    err = errors.New("Group not found")
  }

  return admin, err
}

func NotifyAdmin(company, id string, data []byte) error {
  admin, err := GetAdmin(company, id)
  if err == nil {
    msg := string(data)
    if err := websocket.Message.Send(admin.Connection, msg); err != nil {
      return errors.New("Can't send notify to Admin")
    } else {
      return nil
    }
  } else {
    return errors.New("Admin not found")
  }
}
