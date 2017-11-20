package router

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"db"
  "config"
)

func getConfigurationRouter() RouteList {
	prefix := "/config"
	var routes = RouteList{
		Route{ "LoadConfiguration", "POST",   prefix + "/load", loadConfig, true, },
	}

  return routes
}

func loadConfig(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
  id, err := getObjectId(r)
	if err == nil {
		c := db.NewCollectionSession("admin_db", "Configuration")
		defer c.Close()

		var data config.ConfigData
		err = c.Session.Find(bson.M{"_id": id}).One(&data)
		if err == nil {
      config.SetConfiguration(data);
			responseJsonWithError(w, data, http.StatusOK)
			text = ""
		 }
	}

	responseWithError(w, text, http.StatusInternalServerError)
}
