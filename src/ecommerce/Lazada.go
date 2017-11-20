package ecommerce

import (
  //"fmt"
	"encoding/json"
  "encoding/xml"
	"net/http"
  "time"
  "encoding/hex"
  "crypto/hmac"
  "crypto/sha256"
  "bytes"
	//"os"
  //"io/ioutil"
  "io"
  "errors"
)

var formatLazadaDate string = "2006-01-02T15:04:05-07:00"

func getLazadaTime() string {
  return time.Now().UTC().Format(formatLazadaDate)
}

func (config LazadaConfig) createUrl(method string) string {
  time := getLazadaTime()
  list := make(map[string]string)
  list["Action"] = method
  list["Filter"] = "all"
  list["Format"] = "json"
  list["Timestamp"] = time
  list["UserID"] = config.UserId
  list["Version"] = config.Version
  decoded := []byte(config.Key)
  param, _ := generateParams(list)
  data := []byte(param)
  mac := hmac.New(sha256.New, decoded)
  mac.Write(data)
  expectedMAC := mac.Sum(nil)
  signature := hex.EncodeToString(expectedMAC)
  var buffer bytes.Buffer
  buffer.WriteString(config.Url)
  buffer.WriteString(param)
  buffer.WriteString("&Signature=")
  buffer.WriteString(signature)

  return buffer.String()
}

func (config LazadaConfig) LazadaUploadImage(path string) (string, error) {
  url := config.createUrl("UploadImage")
  response, err := http.Get(path)
  if err != nil {
    return "Cannot get Image URL", err
  }
  defer response.Body.Close()

  buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

  res, err := http.Post(url, "binary/octet-stream", buf)
  if err != nil {
    panic(err)
    return "Cannot Upload Image", err
  }
  defer res.Body.Close()

  decoder := json.NewDecoder(res.Body)
	var result LazadaResImageUpload
	err = decoder.Decode(&result)

  if result.Response.Body.Image.Code != "" {
    return result.Response.Body.Image.Url, nil
  } else {
    return "", errors.New("Image Upload fail")
  }
}

func (config LazadaConfig) LazadaCreateProduct(product LazadaRequest) (string, error) {
  url := config.createUrl("CreateProduct")
  data := new(bytes.Buffer)
  //json.NewEncoder(data).Encode(product)
  xml.NewEncoder(data).Encode(product.Request)

  return config.LazadaPostAction("Create Product", url, data)
}

func (config LazadaConfig) LazadaUpdateProduct(product LazadaRequest) (string, error) {
  url := config.createUrl("UpdateProduct")
  data := new(bytes.Buffer)
  xml.NewEncoder(data).Encode(product.Request)

  return config.LazadaPostAction("Update Product", url, data)
}

func (config LazadaConfig) LazadaPriceQuantity(product LazadaQuantityRequest) (string, error) {
  url := config.createUrl("UpdatePriceQuantity")
  data := new(bytes.Buffer)
  xml.NewEncoder(data).Encode(product.Request)

  return config.LazadaPostAction("Update Price&Qty", url, data)
}

func (config LazadaConfig) LazadaUpdateImage(product LazadaImageRequest) (string, error) {
  url := config.createUrl("SetImages")
  data := new(bytes.Buffer)
  xml.NewEncoder(data).Encode(product.Request)

  return config.LazadaPostAction("Update Images", url, data)
}

func (config LazadaConfig) LazadaPostAction(name, url string, data io.Reader) (string, error) {

  res, err := http.Post(url, "application/xml", data)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var resErr LazadaResError
	err = decoder.Decode(&resErr)

  if resErr.Respond.Header.ErrorCode == 0 {
    return name + " Success", nil
  } else {
    textB, _ := json.Marshal(resErr.Respond.Body)
    text := string(textB)
    return "", errors.New(text)
  }
}

func (config LazadaConfig) LazadaGetCategory() (string, error) {
  url := config.createUrl("GetCategoryTree")

  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var ResErr LazadaResError
	decoder.Decode(&ResErr)

  if ResErr.Respond.Header.ErrorCode != 0 {
    return "Update Product Success", nil
  } else {
    return "", errors.New("Create Product fail")
  }
}

func (config LazadaConfig) LazadaGetCategoryAttribute(category_id int) (string, error) {
  url := config.createUrl("GetCategoryAttributes")

  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var ResErr LazadaResError
	decoder.Decode(&ResErr)

  if ResErr.Respond.Header.ErrorCode != 0 {
    return "Update Product Success", nil
  } else {
    return "", errors.New("Create Product fail")
  }
}
