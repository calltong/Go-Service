package router

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
	"strings"
	"utility"
	"time"
	"fmt"
	"db"
)

func getPageRouter() RouteList {
	prefix := "/page"
	var routes = RouteList{
		Route{ "PageEdit",   "PUT",  prefix + "/{id}/edit", editPage, true, },
		Route{ "PageBuild",  "PUT",  prefix + "/build", buildPage, true, },
		Route{ "PageGetItem","GET",  prefix + "/{page}/{status}", getPageItem, false, },
	}

  return routes
}

func getPageData(w http.ResponseWriter, r *http.Request) (Page, error){
	decoder := json.NewDecoder(r.Body)
	var data Page
	err := decoder.Decode(&data)

	return data, err
}

func editPage(w http.ResponseWriter, r *http.Request) {
	var text string = "Update fail"
	data, err := getPageData(w, r)
	if err == nil {
		access, err := getAccessInfo(r)
		if err == nil {
			if data.Page == "home" {
				data = uploadHomeImage(data, access)
			} else if data.Page == "about_us" {
				data = uploadAboutusImage(data, access)
			}
		}
		text = ""
		editData(w, r, "Page", data)
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func getPageItem(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Page")
		defer c.Close()

		vars := mux.Vars(r)
		page := vars["page"]
		status := vars["status"]
		// get Page
		var data Page
		err = c.Session.Find(bson.M{"page": page, "status": status}).One(&data)
		if err == nil {
			responseJsonWithError(w, data, http.StatusOK)
			text = ""
		}
	}

	responseWithError(w, text, http.StatusInternalServerError)
}

func buildPage(w http.ResponseWriter, r *http.Request) {
	var text string = "Build Page fail"
	access, err := getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Page")
		defer c.Close()

		_, err = c.Session.RemoveAll(bson.M{"status": "backup"})
		if err == nil {
			update := bson.M{"status": "backup"}
			_, err = c.Session.UpdateAll(bson.M{"status": "active"}, bson.M{"$set": update})
			if err == nil {
				var list PageList
				err = c.Session.Find(bson.M{"status": "modify"}).All(&list)

				if err == nil {
					text = ""
					for _, item := range list {
						item.Id = bson.NewObjectId()
						item.Status = "active"
						err = c.Session.Insert(item)
						if err != nil {
							text = "Build and Insert Fail"
							break
						}
					}
				}
		 	}
	 	}
	}
	if text == "" {
		responseWithError(w, "Done", http.StatusOK)
	} else {
		responseWithError(w, text, http.StatusInternalServerError)
	}
}

func uploadHomeImage(page Page, access AccessInfo) Page {
	var data HomeData
	raw, _ := bson.Marshal(page.Data)
	bson.Unmarshal(raw, &data)
	path := fmt.Sprintf("%s/page/%s", access.Project.Folder, page.Id.Hex())
	idea := time.Now().Unix()
	index := 0
	updated := false
	for _, content := range data.List {
		for i, item := range content.Data.List {
			if strings.HasPrefix(item.Preview, "data:image") {
				index += 1
				name := fmt.Sprintf("%d%d.jpg", idea, index)
				temp, _ := utility.UploadImage(item.Preview, path, name)
				npath := fmt.Sprintf("%s/%s", access.Project.Address, temp)
				item.Preview = npath//"http://localhost:8001/resource/seocy/page/5953d02ca5a169c4d25d05a5/15043400811.jpg"
				updated = true
			}
			content.Data.List[i] = item
		}
	}
	if updated {
		var bs bson.M
		raw, _ = bson.Marshal(data)
		bson.Unmarshal(raw, &bs)
		page.Data = bs
	}

	return page;
}

func uploadAboutusImage(page Page, access AccessInfo) Page {
	var data AboutusData
	raw, _ := bson.Marshal(page.Data)
	bson.Unmarshal(raw, &data)
	path := fmt.Sprintf("%s/page/%s", access.Project.Folder, page.Id.Hex())
	idea := time.Now().Unix()
	index := 0
	updated := false
	for i, item := range data.List {
		if strings.HasPrefix(item.Preview, "data:image") {
			index += 1
			name := fmt.Sprintf("%d%d.jpg", idea, index)
			temp, _ := utility.UploadImage(item.Preview, path, name)
			npath := fmt.Sprintf("%s/%s", access.Project.Address, temp)
			item.Preview = npath
			updated = true
		}
		data.List[i] = item
	}

	if updated {
		var bs bson.M
		raw, _ = bson.Marshal(data)
		bson.Unmarshal(raw, &bs)
		page.Data = bs
	}

	return page;
}
