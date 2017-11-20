package admin

import (
	"encoding/json"
	"net/http"
  "github.com/gorilla/mux"
)
// Get Parameter
func getParameter(r *http.Request) (string, string) {
	vars := mux.Vars(r)
	company := vars["company"]
	id := vars["id"]
	return company, id
}

// Get Body
func getAdminData(r *http.Request) (AdminData, error){
	decoder := json.NewDecoder(r.Body)
	var data AdminData
	err := decoder.Decode(&data)

	return data, err
}

func getActiveAdminData(r *http.Request) (ActiveAdminData, error){
	decoder := json.NewDecoder(r.Body)
	var data ActiveAdminData
	err := decoder.Decode(&data)

	return data, err
}
