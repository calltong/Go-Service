package router

type ConnectReq struct {
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  NetUser NetworkData `json:"net_user" bson:"net_user"`
}

type ConnectRes struct {
  SessionId string `json:"session_id" bson:"session_id"`
  AdminId string `json:"admin_id" bson:"admin_id"`
  NetAdmin NetworkData `json:"net_admin" bson:"net_admin"`
}

type ReconnectReq struct {
  Company string `json:"company" bson:"company"`
  UserId string `json:"user_id" bson:"user_id"`
  SessionId string `json:"session_id" bson:"session_id"`
  AdminId string `json:"admin_id" bson:"admin_id"`
  NetUser NetworkData `json:"net_user" bson:"net_user"`
}
