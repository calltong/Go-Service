package user

type TestData struct {
  Method string `json:"method" bson:"method"`
  Type string `json:"type" bson:"type"`
  UserId string `json:"user_id" bson:"user_id"`
}
