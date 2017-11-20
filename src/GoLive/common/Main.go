package common

import (
	"encoding/json"
	"net/http"
)

func GetLoginData(r *http.Request) (LoginData, error){
	decoder := json.NewDecoder(r.Body)
	var data LoginData
	err := decoder.Decode(&data)

	return data, err
}
