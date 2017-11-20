package ecommerce

import (
  //"fmt"
	"encoding/json"
  "encoding/xml"
	"net/http"
  "time"
  "bytes"
  "io"
  "io/ioutil"
  "errors"
  "os"
)

var formatStreetDate string = "2006-01-02"

func getStreetTime() string {
  return time.Now().UTC().Format(formatStreetDate)
}

func (config StreetConfig) streetCreateUrl(method string) string {

  var buffer bytes.Buffer
  buffer.WriteString(config.Url)
  buffer.WriteString(method)

  return buffer.String()
}

func (config StreetConfig) StreetUploadImage(path string) (string, error) {
  url := config.streetCreateUrl("UploadImage")
  file, err := os.Open(path)
  if err != nil {
    panic(err)
  }
  defer file.Close()

  res, err := http.Post(url, "binary/octet-stream", file)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()

  decoder := json.NewDecoder(res.Body)
	var result LazadaResImageUpload
	err = decoder.Decode(&result)
  if err == nil {
    return result.Response.Body.Image.Url, err
  } else {
    body, _ := ioutil.ReadAll(res.Body)
    return string(body), errors.New("Image Upload fail")
  }
}

func (config StreetConfig) StreetCreateProduct(product StreetProductRequest) (string, error) {
  url := config.streetCreateUrl("prodservices/product")
  data := new(bytes.Buffer)
  xml.NewEncoder(data).Encode(product.Product)

  return config.StreetPostAction("Create Product", url, data)
}

func (config StreetConfig) StreetPostAction(name, url string, data io.Reader) (string, error) {
  client := &http.Client { }
  req, err := http.NewRequest("POST", url, data)
  req.Header.Set("Content-Type", "application/xml")
  req.Header.Set("openapikey", config.Key)
  res, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  decoder := json.NewDecoder(res.Body)
  var resMsg StreetRespond
	err = decoder.Decode(&resMsg)

  if resMsg.Product.Code == 0 {
    return name + " Success", nil
  } else {
    textB, _ := json.Marshal(resMsg)
    text := string(textB)
    return text, errors.New(name + " fail")
  }
}
