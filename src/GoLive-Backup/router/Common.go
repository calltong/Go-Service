package router

import (
  "encoding/json"
  "gopkg.in/mgo.v2/bson"
)

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
