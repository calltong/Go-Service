package router

import (
  "gopkg.in/mgo.v2/bson"
  "net/http"
  "authentication"
)

type AccessInfo struct {
  User bson.ObjectId `json:"user"`
  Project ProjectGeneral `json:"project"`
}

func getAccessInfo(r *http.Request) (AccessInfo, error) {
  var data AccessInfo
  token := r.Header.Get("authorization")
  decode, err := authentication.DecodeToken(token);
  data.User = decode.User
  data.Project.Folder = decode.Project.Folder
  data.Project.Database = decode.Project.Database
  data.Project.Address = decode.Project.Address

  return data, err
}

func getToken(id bson.ObjectId, project ProjectGeneral) string {
  var data authentication.ProjectGeneral
  data.Folder = project.Folder
  data.Database = project.Database
  data.Address = project.Address

  return authentication.CreateToken(id, data);
}
