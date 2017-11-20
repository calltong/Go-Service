package router

import (
	"encoding/json"
	"time"
	//"fmt"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"db"
	"authentication"
	"config"
)

func getUserRouter() RouteList {
	prefix := "/user"
	var routes = RouteList {
    Route{ "UserLogin",  "PUT", prefix + "/login", getUserLoginItem, false, },
		Route{ "VerifyToken","PUT", prefix + "/tokenverify", verifyToken, false, },
		Route{ "UserGet", "PUT", prefix + "/get", getUserItem, true, },
		Route{ "UserToken", "GET", prefix + "/gettoken", getTokenByOrigin, false, },
	}

  return routes
}

func getUserData(w http.ResponseWriter, r *http.Request) (User, error){
	decoder := json.NewDecoder(r.Body)
	var data User
	err := decoder.Decode(&data)

	return data, err
}

func getUserLoginData(w http.ResponseWriter, r *http.Request) (UserLogin, error){
	decoder := json.NewDecoder(r.Body)
	var data UserLogin
	err := decoder.Decode(&data)

	return data, err
}

func getProject(dbname string, id bson.ObjectId) (Project, error) {
	c := db.NewCollectionSession(dbname, "Project")
	defer c.Close()

	var data Project
	err := c.Session.Find(bson.M{"_id": id}).One(&data);
	return data, err
}

func getUserLoginItem(w http.ResponseWriter, r *http.Request) {
	var text string = "Input Data Incorrect"
	data, err := getUserLoginData(w, r)
	if err == nil {
		conf := config.GetConfiguration()
		dbname := conf.Database.DbName
		c := db.NewCollectionSession(dbname, "User")
		defer c.Close()

		var user User
		err = c.Session.Find(bson.M{"username": data.Username, "password": data.Password}).One(&user);
		if err != nil {
			text = "Username or Password Incorrect"
		} else {
			find := user.ProjectList[0];
			project, err := getProject(dbname, find.Id)
			if err != nil {
				text = "User have no Project"
			} else {
				var token TokenData
				token.Token = getToken(user.Id, project.General)
				responseJSON(w, token, http.StatusOK)
				text = ""
			}
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func getUserItem(w http.ResponseWriter, r *http.Request) {
	var text string = "Input Data Incorrect"
	access, err := getAccessInfo(r)
	if err == nil {
		conf := config.GetConfiguration()
		dbname := conf.Database.DbName
		c := db.NewCollectionSession(dbname, "User")
		defer c.Close()

		var user User
		err = c.Session.Find(bson.M{"_id": access.User}).One(&user);
		if err != nil {
			text = "Token Incorrect"
		} else {
			user.Username = ""
			user.Password = ""
			responseJSON(w, user, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func verifyToken(w http.ResponseWriter, r *http.Request) {
	var text string = "Token Incorrect"
	decoder := json.NewDecoder(r.Body)
	var data TokenData
	err := decoder.Decode(&data)
	if err == nil {
		claim, err := authentication.DecodeToken(data.Token)
 		if err == nil {
			current := time.Now().Unix()
			if current < claim.ExpiresAt {
				text = ""
				responseText(w, "passed", http.StatusOK);
			}
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}

func getTokenByOrigin(w http.ResponseWriter, r *http.Request) {
	var text string = "Not Found Data"
	conf := config.GetConfiguration()
	dbname := conf.Database.DbName
	c := db.NewCollectionSession(dbname, "User")
	defer c.Close()

	origin := r.Header.Get("Origin")
	var user User
	err := c.Session.Find(bson.M{"address_list": origin}).One(&user);
	if err == nil {
		find := user.ProjectList[0];
		project, err := getProject(dbname, find.Id)
		if err != nil {
			text = "User have no Project"
		} else {
			var token TokenData
			token.Token = getToken(user.Id, project.General)
			responseJSON(w, token, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusBadRequest)
}
