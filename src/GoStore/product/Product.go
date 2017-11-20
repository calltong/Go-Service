package product

import (
	"encoding/json"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"time"
	"strings"
	"fmt"
	"utility"
	"common/router"
	"db"
)

func GetRouter() router.RouteList {
	prefix := "/product"
	var routes = router.RouteList {
		router.Route{ "ProductCreate", "POST",   prefix + "/create", createProduct, false, },
		router.Route{ "ProductDelete", "DELETE", prefix + "/{id}/delete", deleteProduct, false, },
		router.Route{ "ProductEdit",   "PUT",  prefix + "/{id}/edit", editProduct, false, },
		router.Route{ "ProductGetList","GET",  prefix, getProductList, false, },
		router.Route{ "ProductGetItem","GET",  prefix + "/{id}/get", getProductItem, false, },
		router.Route{ "ProductCount",  "GET",  prefix + "/count", getProductCount, false, },
	}

  return routes
}

func getData(w http.ResponseWriter, r *http.Request) (Product, error){
	decoder := json.NewDecoder(r.Body)
	var data Product
	err := decoder.Decode(&data)

	return data, err
}

func Create(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not created"
	var access AccessInfo
  var data Product
	var err error = nil
	data, err = getProductData(w, r)
	if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			data.Id = bson.NewObjectId()
			data = setupImage(data, access)
			data.LastUpdate = time.Now().Unix()
			createData(w, r, "Product", data)
			text = ""
		}
	}

	router.ResponseWithError(w, text, http.StatusBadRequest)
	//LogServerError(access, "Router/createProduct", err)
}

func Remove(w http.ResponseWriter, r *http.Request) {
	deleteData(w, r, "Product")
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not updated"
	var access AccessInfo
  var data Product
	var err error = nil
	data, err = getProductData(w, r)
	if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			data = setupImage(data, access)
			data.LastUpdate = time.Now().Unix()
			editData(w, r, "Product", data)
			text = ""
		}
	}

	router.ResponseWithError(w, text, http.StatusBadRequest)
	//LogServerError(access, "Router/editProduct", err)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	var access AccessInfo
  var list ProductList
	var err error = nil
	access, err = getAccessInfo(r)
	if err == nil {
		queries := r.URL.Query()
		typ := queries.Get("type")
		limit := convertToInt(queries.Get("limit"), 40)
		page := convertToInt(queries.Get("page"), 1) - 1

		var condition bson.M
		if typ != "" {
			condition = bson.M{"type_id": bson.ObjectIdHex(typ)}
		}

		list, err = getProductListWithCondition(access, condition, page * limit, limit)
		if err == nil {
			router.ResponseJsonWithError(w, list, http.StatusOK)
			text = ""
		}
	}
	router.ResponseWithError(w, text, http.StatusInternalServerError)
	//LogServerError(access, "Router/getProductList", err)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	var data Product
	GetDataItem(w, r, "Product", data)
}

func uploadImage(idea int64, index int, img, path string, access AccessInfo) string {
	if strings.HasPrefix(img, "data:image") {
		name := fmt.Sprintf("i%d%d.jpg", idea, index)
		temp, _ := utility.UploadImage(img, path, name)
		return fmt.Sprintf("%s/%s", access.Project.Address, temp)
	} else {
		return img
	}
}

func setupImage(data Product, access AccessInfo) Product {
	path := fmt.Sprintf("%s/product/%s", access.Project.Folder, data.Id.Hex())
	utility.CreatePath(path)
	idea := time.Now().Unix()
	size := len(data.VariantList)
	variants := make([]Variant, size)
	for index, variant := range data.VariantList {
		size = len(variant.ImageList)
		list := make([]string, size)
		for i, img := range variant.ImageList {
			list[i] = uploadImage(idea, i, img, path, access)
		}
		variant.ImageList = list

		sqSize := len(variant.ImageSqList)
		list = make([]string, sqSize)
		for i, img := range variant.ImageSqList {
			list[i] = uploadImage(idea, i + size, img, path, access)
		}
		variant.ImageSqList = list
		variants[index] = variant
	}

	data.VariantList = variants
	for _, variant := range data.VariantList {
		if len(variant.ImageList) > 0 {
			data.Image = variant.ImageList[0]
			break
		}
	}

	return data
}


func getProductListWithCondition(access AccessInfo, condition bson.M, skip, limit int) (ProductList, error){
	c := db.NewCollectionSession(access.Project.Database, "Product")
	defer c.Close()
	// get list
	var list ProductList
	err := c.Session.Find(condition).Sort("last_update").Skip(skip).Limit(limit).All(&list)

	return list, err
}

func CountItem(w http.ResponseWriter, r *http.Request) {
	var text string = "Data not found"
	var access AccessInfo
  var num int
	var err error = nil
	access, err = getAccessInfo(r)
	if err == nil {
		c := db.NewCollectionSession(access.Project.Database, "Product")
		defer c.Close()
		queries := r.URL.Query()
		typ := queries.Get("type")

		var condition bson.M
		if typ != "" {
			condition = bson.M{"type_id": bson.ObjectIdHex(typ)}
		}

		// get list
	  num, err = c.Session.Find(condition).Count()
		if err == nil {
			data := Total {Value:num}
			responseJsonWithError(w, data, http.StatusOK)
			text = ""
		}
	}
	router.ResponseWithError(w, text, http.StatusInternalServerError)
	//LogServerError(access, "Router/getProductCount", err)
}
