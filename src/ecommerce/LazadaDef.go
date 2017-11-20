package ecommerce

import (
  "gopkg.in/mgo.v2/bson"
)

type LazadaProductAtt struct {
  Name string `json:"name" bson:"name" xml:"name"`
  ShortDescription string `json:"short_description" bson:"short_description" xml:"short_description"`
  Description string `json:"description" bson:"description" xml:"description"`
  Brand string `json:"brand" bson:"brand" xml:"brand"`
  Model string `json:"model" bson:"model" xml:"model"`
  Color string `json:"color_family" bson:"color_family" xml:"color_family"`
  BraType string `json:"bras_types" bson:"bras_types" xml:"bras_types"`
  SleeperType string `json:"sleep_lounge_styles" bson:"sleep_lounge_styles"`
  WarrantyType string `json:"warranty_type" bson:"warranty_type" xml:"warranty_type"`
  Warranty string `json:"warranty" bson:"warranty" xml:"warranty"`
  NameEn string `json:"name_en" bson:"name_en" xml:"name_en"`
  DescriptionEn string `json:"description_en" bson:"description_en" xml:"description_en"`
}

type LazadaSku struct {
  SellerSku string `json:"SellerSku" bson:"SellerSku" xml:"SellerSku"`
  Size string `json:"size" bson:"size" xml:"size"`
  Quantity int `json:"quantity" bson:"quantity" xml:"quantity"`
  Price int `json:"price" bson:"price" xml:"price"`
  SpecialPrice int `json:"special_price" bson:"special_price" xml:"special_price"`
  SpecialToDate string `json:"special_to_date" bson:"special_to_date" xml:"special_to_date"`
  SpecialFromDate string `json:"special_from_date" bson:"special_from_date" xml:"special_from_date"`
  PackageLength string `json:"package_length" bson:"package_length" xml:"package_length"`
  PackageWeight string `json:"package_weight" bson:"product_weight" xml:"package_weight"`
  PackageHeight string `json:"package_height" bson:"package_height" xml:"package_height"`
  PackageWidth string `json:"package_width" bson:"package_width" xml:"package_width"`
  PackageContent string `json:"package_content" bson:"package_content" xml:"package_content"`
  ImageList []string `json:"Images" bson:"Images" xml:"Images>Image"`
}

type LazadaProductContent struct {
  PrimaryCategory int `json:"PrimaryCategory" bson:"PrimaryCategory" xml:"PrimaryCategory"`
  Attribute LazadaProductAtt `json:"Attributes" bson:"Attributes" xml:"Attributes"`
  SkuList []LazadaSku `json:"Skus" bson:"Skus" xml:"Skus>Sku"`
}

type LazadaProduct struct {
  Product LazadaProductContent `json:"Product" bson:"Product" xml:"Product"`
}

type LazadaRequest struct {
  Request LazadaProduct `json:"Request" bson:"Request" xml:"Request"`
}

/// Only Quantity data
/// ----------------------
type LazadaQuantity struct {
  SellerSku string `json:"SellerSku" bson:"SellerSku" xml:"SellerSku"`
  Quantity int `json:"quantity" bson:"quantity" xml:"Quantity"`
  Price int `json:"price" bson:"price" xml:"Price"`
  SpecialPrice int `json:"special_price" bson:"special_price" xml:"SalePrice"`
  SpecialToDate string `json:"special_to_date" bson:"special_to_date" xml:"SaleEndDate"`
  SpecialFromDate string `json:"special_from_date" bson:"special_from_date" xml:"SaleStartDate"`
}

type LazadaQuantityContent struct {
  SkuList []LazadaQuantity `json:"Skus" bson:"Skus" xml:"Skus>Sku"`
}

type LazadaQuantityProduct struct {
  Product LazadaQuantityContent `json:"Product" bson:"Product" xml:"Product"`
}

type LazadaQuantityRequest struct {
  Request LazadaQuantityProduct `json:"Request" bson:"Request" xml:"Request"`
}

/// Only Image list
/// ----------------------
type LazadaImage struct {
  SellerSku string `json:"SellerSku" bson:"SellerSku" xml:"SellerSku"`
  ImageList []string `json:"Images" bson:"Images" xml:"Images>Image"`
}

type LazadaImageContent struct {
  SkuList []LazadaImage `json:"Skus" bson:"Skus" xml:"Skus>Sku"`
}

type LazadaImageProduct struct {
  Product LazadaImageContent `json:"Product" bson:"Product" xml:"Product"`
}

type LazadaImageRequest struct {
  Request LazadaImageProduct `json:"Request" bson:"Request" xml:"Request"`
}

/// Lazada Config
/// ---------------------------

type LazadaConfig struct {
  Url string `json:"url" bson:"url"`
  UserId string `json:"user_id" bson:"user_id"`
  Key string `json:"key" bson:"key"`
  Version string `json:"version" bson:"version"`
}

// Lazada Respond error
type LazadaErrorHeader struct {
  ErrorCode int `json:"ErrorCode" bson:"ErrorCode"`
  ErrorMessage string `json:"ErrorMessage" bson:"ErrorMessage"`
  ErrorType string `json:"ErrorType" bson:"ErrorType"`
  RequestAction string `json:"RequestAction" bson:"RequestAction"`
  RequestId string `json:"RequestId" bson:"RequestId"`
}

type LazadaMsgError struct {
  Header LazadaErrorHeader `json:"Head" bson:"Head"`
  Body bson.M `json:"Body" bson:"Body"`
}

type LazadaResError struct {
  Respond LazadaMsgError `json:"ErrorResponse" bson:"ErrorResponse"`
}

// Lazada Respond Message
type LazadaHeader struct {
  RequestId string `json:"RequestId" bson:"RequestId"`
  RequestAction string `json:"RequestAction" bson:"RequestAction"`
  ResponseType string `json:"ResponseType" bson:"ResponseType"`
  Timestamp string `json:"Timestamp" bson:"Timestamp"`
  isES bool `json:"isES" bson:"isES"`
}

type LazadaMsgRespond struct {
  Header LazadaHeader `json:"Head" bson:"Head"`
}

type LazadaRespond struct {
  Response LazadaMsgRespond `json:"SuccessResponse" bson:"SuccessResponse"`
}

// Image Respond
type LazadaBodyImageUpload struct {
  Image LazadaImageUpload `json:"Image" bson:"Image"`
}

type LazadaImageUpload struct {
  Url string `json:"Url" bson:"Url"`
  Code string `json:"Code" bson:"Code"`
}

type LazadaMsgImageUpload struct {
  Header LazadaHeader `json:"head" bson:"head"`
  Body LazadaBodyImageUpload `json:"body" bson:"body"`
}

type LazadaResImageUpload struct {
  Response LazadaMsgImageUpload `json:"SuccessResponse" bson:"SuccessResponse"`
}
