package router

import (
  "gopkg.in/mgo.v2/bson"
  "fmt"
	"time"
  "db"
)

func (data LogValue) Add(access AccessInfo, system string) {
  year := time.Now().UTC().Format("2016")

  c := db.NewCollectionSession(access.Project.Database, "Log")
  defer c.Close()

  c.Session.Upsert(bson.M{"year": year, "system": system}, bson.M{"$push": bson.M{"list": data}})
}

func LogErrorMessage(access AccessInfo, system, source, message string) {
  date := time.Now()
  year := fmt.Sprintf("%04v-%02d", date.Year(), date.Month())
  c := db.NewCollectionSession(access.Project.Database, "Log")
  defer c.Close()

  var data LogValue
  data.Date = date.Unix()
  data.Type = "error"
	data.Source = source
	data.Message = message

  c.Session.Upsert(bson.M{"year": year, "system": system}, bson.M{"$push": bson.M{"list": data}})
}

func LogError(access AccessInfo, system, source string, err error) {
  if err != nil {
    LogErrorMessage(access, system, source, err.Error())
  }
}

func LogServerError(access AccessInfo, source string, err error) {
  if err != nil {
    LogErrorMessage(access, "server", source, err.Error())
  }
}
