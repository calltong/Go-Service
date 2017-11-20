package router

import (
  //"fmt"
	//"gopkg.in/mgo.v2/bson"
	"net/http"
  //"bytes"
  "ecommerce"
  "time"
	"errors"
)

func getStreetRouter() RouteList {
	prefix := "/ecommerce/street/"
	var routes = RouteList{
    Route{ "STREET_UPDATE_PRODUCT",  "PUT",  prefix + "{id}/update", lazadaUpdateProduct, false, },
		//Route{ "STREET_UPDATE_QUANTITY", "PUT",  prefix + "{id}/quantity", lazadaUpdateQuantity, false, },
		//Route{ "STREET_UPDATE_IMAGE",  "PUT",  prefix + "{id}/image", lazadaUpdateImage, false, },
		//Route{ "STREET_CHECK_HAVE",  "GET",  prefix + "{id}/check", lazadaCheckProduct, false, },
	}

  return routes
}

func (product Product) convertStreetProduct(access AccessInfo, setting SettingStore, images []string) ecommerce.StreetProductRequest {
  var data ecommerce.StreetProduct

  begin := time.Now().UTC()
  end := begin.AddDate(0, setting.Street.Product.DurationPromotion, 0)
  //proType, _ := getProductType(product.TypeId, access)

	//data.Name = product.Name
	data.SaleType = "01"
  data.Category = 0//proType.Ecommerce.Street.Category
	data.ServiceType = "01"
	data.Condition = "01"
	data.Weight = "0.3"

	data.HtmlDetail = ""
	data.SellTermUseYn = "N"
	data.StartDate = begin.Format("2006-01-02")
	data.EndDate = end.Format("2006-01-02")
	data.Price = product.Price
	data.ProductStock = 5
	data.AmountOfPurchase = 2
	data.Contact = ""
	data.DeliveryCondition = "30139"
	data.ReturnExchange = ""
	data.VatCode = "01"

	for index, item := range images {
		switch index {
		case 0:
			data.Image01 = item
			break;
		case 1:
			data.Image02 = item
			break;
		case 2:
			data.Image03 = item
			break;
		case 3:
			data.Image04 = item
			break;
		case 4:
			data.Image05 = item
			break;
		case 5:
			data.Image06 = item
			break;
		case 6:
			data.Image07 = item
			break;
		case 7:
			data.Image08 = item
			break;
		}
	}
	/*
  var list []ecommerce.StreetSku
  for _, item := range product.StockList {
    var sku ecommerce.StreetSku
    sku.Name = item.Size.Code
    sku.Quantity = item.Quantity
    sku.Status = "Y"

    list = append(list, sku)
  }

  data.SkuList = list
	*/
  var content ecommerce.StreetProductRequest
  content.Product = data
  return content
}
/*
func (product Product) convertStreetQuantity(lazada ecommerce.LazadaRequest, setting SettingData) (ecommerce.LazadaRequest, ecommerce.LazadaQuantityRequest, error) {
  var data ecommerce.LazadaQuantityContent
  var content ecommerce.LazadaQuantityRequest

  begin := time.Now().UTC()
  end := begin.AddDate(0, setting.Lazada.Product.DurationPromotion, 0)

  var list []ecommerce.LazadaQuantity
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
  content.Request.Product = data
  return lazada, content, nil
}

func (product Product) convertStreetImage(lazada ecommerce.LazadaRequest, images []string) (ecommerce.LazadaRequest, ecommerce.LazadaImageRequest, error) {
  var data ecommerce.LazadaImageContent
  var content ecommerce.LazadaImageRequest

  var list []ecommerce.LazadaImage
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
  content.Request.Product = data
  return lazada, content, nil
}
*/

func getStreetConfig(setting SettingStore) ecommerce.StreetConfig {
  var config ecommerce.StreetConfig
  config.Url = setting.Street.User.Url
  config.Key = setting.Street.User.Key
  config.Version = setting.Street.User.Version

  return config
}

func streetUploadImage(list ImageDataList, access AccessInfo, config ecommerce.StreetConfig) ([]string, error){
	var images []string
	var err error
	for _, img := range list {
		images = append(images, img.Data)
		//path := replacePath(img.Data, access)
		//npath, err := config.StreetUploadImage(path)
		//if err == nil {
		//	images = append(images, npath)
		//} else {
		//	break
		//}
	}
	if len(images) == 0 {
		return images, errors.New("No Images")
	} else {
		return images, err
	}
}

func streetUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var text string = "Data cannot updated"
	/*id, access, product, err := getProductByRequest(r)
	if err == nil {
		setting, _ := getSettingProject(access)
		config := getStreetConfig(setting)
		streetProduct, err := getStreetProduct(id, access)
		if err == nil { // create product on lazada
      images := streetProduct.Content.Product.SkuList[0].ImageList
			lazada := product.convertLazadaProduct(access, setting, images)
			_, err = config.StreetUpdateProduct(lazada)

			if err == nil {
        text = ""
				updateStreetProduct(access, streetProduct)
				responseText(w, "The product updated", http.StatusOK)
			}

		} else { // update product on lazada
			images, err := streetUploadImage(product.ImageSquareList, access, config)

			if err == nil {
				data := product.convertStreetProduct(access, setting, images)
				_, err = config.StreetCreateProduct(data)

				if err == nil {
					text = ""
					var nProduct StreetProductItem
					nProduct.Id = id
					nProduct.Content = data
					insertStreetProduct(access, nProduct)
					responseText(w, "The product created", http.StatusOK)
				}
			}
		}*/
		responseWithError(w, text, http.StatusBadRequest)
	}
