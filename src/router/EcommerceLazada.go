package router

import (
  //"fmt"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
  //"bytes"
  "ecommerce"
  //"time"
	"errors"
)

func getLazadaRouter() RouteList {
	prefix := "/ecommerce/lazada/"
	var routes = RouteList{
    Route{ "LAZADA_UPDATE_PRODUCT",  "PUT",  prefix + "{id}/update", lazadaUpdateProduct, false, },
		Route{ "LAZADA_UPDATE_QUANTITY", "PUT",  prefix + "{id}/quantity", lazadaUpdateQuantity, false, },
		Route{ "LAZADA_UPDATE_IMAGE",  "PUT",  prefix + "{id}/image", lazadaUpdateImage, false, },
		Route{ "LAZADA_CHECK_HAVE",  "GET",  prefix + "{id}/check", lazadaCheckProduct, false, },
	}

  return routes
}

func (product Product) convertLazadaProduct(access AccessInfo, setting SettingStore, images []string) ecommerce.LazadaRequest {
  var data ecommerce.LazadaProductContent

  color := ""
  //begin := time.Now().UTC()
  //end := begin.AddDate(0, setting.Lazada.Product.DurationPromotion, 0)
  proType, _ := getProductType(product.TypeId, access)

  data.PrimaryCategory = proType.Ecommerce.Lazada.CategoryId
  data.Attribute.Name = product.Content.Main.Name
  data.Attribute.ShortDescription = product.Content.Main.Description
  data.Attribute.Description = product.Content.Main.Description
  data.Attribute.NameEn = product.Content.Main.Name
  data.Attribute.DescriptionEn = product.Content.Main.Description
	data.Attribute.Color = color

  data.Attribute.Brand = setting.Lazada.Product.Brand
  data.Attribute.BraType = setting.Lazada.Product.BraType
	data.Attribute.SleeperType = setting.Lazada.Product.SleeperType
  data.Attribute.WarrantyType = setting.Lazada.Product.WarrantyType
  data.Attribute.Warranty = setting.Lazada.Product.Warranty
  data.Attribute.Model = proType.Ecommerce.Lazada.Model
  //var list []ecommerce.LazadaSku
	/*
  for _, item := range product.StockList {
    var sku ecommerce.LazadaSku
    sku.SellerSku = product.Code + "-" + item.Size.Code
    sku.Size = item.Size.Ecommerce.Lazada.Code
    sku.Quantity = item.Quantity
    sku.Price = product.Price
    sku.SpecialPrice = product.SalePrice
    sku.SpecialFromDate = begin.Format("2006-01-02")
    sku.SpecialToDate = end.Format("2006-01-02")
    sku.PackageLength = setting.Lazada.Product.PackageLength
    sku.PackageWeight = setting.Lazada.Product.PackageWeight
    sku.PackageHeight = setting.Lazada.Product.PackageHeight
    sku.PackageWidth = setting.Lazada.Product.PackageWidth
    sku.PackageContent = product.Information.PackageContent
    sku.ImageList = images
    list = append(list, sku)
  }

  data.SkuList = list
	*/
  var content ecommerce.LazadaRequest
  content.Request.Product = data
  return content
}

func (product Product) convertLazadaQuantity(lazada ecommerce.LazadaRequest, setting SettingStore) (ecommerce.LazadaRequest, ecommerce.LazadaQuantityRequest, error) {
  var data ecommerce.LazadaQuantityContent
  var content ecommerce.LazadaQuantityRequest

  //begin := time.Now().UTC()
  //end := begin.AddDate(0, setting.Lazada.Product.DurationPromotion, 0)

  //var list []ecommerce.LazadaQuantity
	/*
  for index, item := range product.StockList {
    var sku ecommerce.LazadaQuantity
    sku.SellerSku = product.Code + "-" + item.Size.Code
    sku.Quantity = item.Quantity
    sku.Price = product.Price
    sku.SpecialPrice = product.SalePrice
    sku.SpecialFromDate = begin.Format("2006-01-02")
    sku.SpecialToDate = end.Format("2006-01-02")
    list = append(list, sku)

    lazadaSku := lazada.Request.Product.SkuList[index]
    if lazadaSku.SellerSku == sku.SellerSku {
      lazadaSku.Quantity = sku.Quantity
      lazadaSku.Price = sku.Price
      lazadaSku.SpecialPrice = sku.SpecialPrice
      lazadaSku.SpecialFromDate = sku.SpecialFromDate
      lazadaSku.SpecialToDate = sku.SpecialToDate
    } else {
      return lazada, content, errors.New("SKU not found")
    }
  }

  data.SkuList = list
	*/
  content.Request.Product = data
  return lazada, content, nil
}

func (product Product) convertLazadaImage(lazada ecommerce.LazadaRequest, images []string) (ecommerce.LazadaRequest, ecommerce.LazadaImageRequest, error) {
  var data ecommerce.LazadaImageContent
  var content ecommerce.LazadaImageRequest

  //var list []ecommerce.LazadaImage
	/*
  for index, item := range product.StockList {
    var sku ecommerce.LazadaImage
    sku.SellerSku = product.Code + "-" + item.Size.Code
    sku.ImageList = images
    list = append(list, sku)

    lazadaSku := lazada.Request.Product.SkuList[index]
    if lazadaSku.SellerSku == sku.SellerSku {
      lazadaSku.ImageList = sku.ImageList
    } else {
      return lazada, content, errors.New("SKU not found")
    }
  }

  data.SkuList = list
	*/
  content.Request.Product = data
  return lazada, content, nil
}

func getLazadaConfig(setting SettingStore) ecommerce.LazadaConfig {
  var config ecommerce.LazadaConfig
  config.Url = setting.Lazada.User.Url
  config.UserId = setting.Lazada.User.Id
  config.Key = setting.Lazada.User.Key
  config.Version = setting.Lazada.User.Version

  return config
}

func lazadaUploadImage(list ImageDataList, access AccessInfo, config ecommerce.LazadaConfig) ([]string, error){
	var images []string
	var err error = nil
	for _, img := range list {
		path, err := config.LazadaUploadImage(img.Data)
		if err == nil {
			images = append(images, path)
		} else {
			break
		}
	}
	if len(images) == 0 {
		return images, errors.New("No Images")
	} else {
		return images, err
	}
}

func lazadaUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var text string = "Data cannot updated"
	/*
  var id bson.ObjectId
  var access AccessInfo
  var product Product
  var lazadaProduct LazadaProductItem
  var images []string
  var err error = nil
	id, access, product, err = getProductByRequest(r)
	if err == nil {
    var setting SettingStore
		setting, _ = getSettingProject(access)
		config := getLazadaConfig(setting)
		lazadaProduct, err = getLazadaProduct(id, access)
		if err == nil { // update product on lazada
			images = lazadaProduct.Content.Request.Product.SkuList[0].ImageList
			lazada := product.convertLazadaProduct(access, setting, images)
			_, err = config.LazadaUpdateProduct(lazada)
			if err == nil {
        text = ""
				var nProduct LazadaProductItem
				nProduct.Id = id
				nProduct.Content = lazada
				updateLazadaProduct(access, nProduct)
				responseText(w, "The product updated", http.StatusOK)
			}
		} else { // create product on lazada
			images, err = lazadaUploadImage(product.ImageSquareList, access, config)
			if err == nil {
				lazada := product.convertLazadaProduct(access, setting, images)
				_, err = config.LazadaCreateProduct(lazada)
				if err == nil {
					text = ""
					var nProduct LazadaProductItem
					nProduct.Id = id
					nProduct.Content = lazada
					insertLazadaProduct(access, nProduct)
					responseText(w, "The product created", http.StatusOK)
				}
			}
		}
	}
	*/
  responseWithError(w, text, http.StatusBadRequest)
  //LogServerError(access, "Router/lazadaUpdateProduct", err)
}

func lazadaUpdateQuantity(w http.ResponseWriter, r *http.Request) {
	var text string = "SKU Data cannot updated"
	/*
  var id bson.ObjectId
  var access AccessInfo
  var product Product
  var lazadaProduct LazadaProductItem
  var err error = nil
	id, access, product, err = getProductByRequest(r)
	if err == nil {
		lazadaProduct, err = getLazadaProduct(id, access)
		if err != nil {
			text = "The product not found"
		} else {
			setting, _ := getSettingProject(access)
			config := getLazadaConfig(setting)
      var lazada ecommerce.LazadaQuantityRequest
			lazadaProduct.Content, lazada, err = product.convertLazadaQuantity(lazadaProduct.Content, setting)
      if err == nil {
        text, err = config.LazadaPriceQuantity(lazada)
  			if err == nil {
  				text = ""
  				updateLazadaProduct(access, lazadaProduct)
  				responseText(w, "The product updated", http.StatusOK)
  			}
      } else {
        text = "SKU List not same"
      }
		}
	}
	*/
  responseWithError(w, text, http.StatusBadRequest)
  //LogServerError(access, "Router/lazadaUpdateQuantity", err)
}

func lazadaUpdateImage(w http.ResponseWriter, r *http.Request) {
	var text string = "Image cannot updated"
	/*
  var id bson.ObjectId
  var access AccessInfo
  var product Product
  var lazadaProduct LazadaProductItem
  var images []string
  var err error = nil
	id, access, product, err = getProductByRequest(r)
	if err == nil {
		lazadaProduct, err = getLazadaProduct(id, access)
		if err != nil {
			text = "The product not found"
		} else {
			setting, _ := getSettingProject(access)
			config := getLazadaConfig(setting)
			images, err = lazadaUploadImage(product.ImageSquareList, access, config)

			if err == nil {
        var lazada ecommerce.LazadaImageRequest
				lazadaProduct.Content, lazada, err = product.convertLazadaImage(lazadaProduct.Content, images)
        if err == nil {
          _, err = config.LazadaUpdateImage(lazada)

  				if err == nil {
  					text = ""
            updateLazadaProduct(access, lazadaProduct)
  					responseText(w, "The product image updated", http.StatusOK)
  				}
        } else {
          text = "SKU List not same"
        }
			}
		}
	}
	*/
  responseWithError(w, text, http.StatusBadRequest)
  //LogServerError(access, "Router/lazadaUpdateImage", err)
}

func lazadaCheckProduct(w http.ResponseWriter, r *http.Request) {
	var text string = "The product not found"
	var access AccessInfo
	id, err := getObjectId(r)
  if err == nil {
		access, err = getAccessInfo(r)
		if err == nil {
			_, err := getLazadaProduct(id, access)
			if err == nil {
				text = ""
				responseText(w, "The product already register", http.StatusOK)
			}
		}
	}
	responseWithError(w, text, http.StatusNotFound)
}
