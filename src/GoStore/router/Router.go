package router

import (
  "net/http"
	"github.com/gorilla/mux"
  "errors"
  "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "common/router"
  "GoStore/product"
)

func CreateAllRouter() *mux.Router {
	routers := mux.NewRouter().StrictSlash(true)
  //addRouter(router, getProductSizeRouter())
  //addRouter(router, getProductTypeRouter())
  router.addRouter(router, product.GetRouter())
  //addRouter(router, getColorRouter())
  //addRouter(router, getPageRouter())
  //addRouter(router, getCustomerRouter())
  //addRouter(router, getUserRouter())
  //addRouter(router, getOrderRouter())
  //addRouter(router, getConfigurationRouter())
  //addRouter(router, getSettingProjectRouter())
  //addRouter(router, getLazadaRouter())

	return routers
}
