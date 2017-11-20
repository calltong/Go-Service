package ecommerce

type StreetSku struct {
  Name string `json:"addPrdGrpNm" bson:"addPrdGrpNm" xml:"addPrdGrpNm"`
  Quantity int `json:"compPrdQty" bson:"compPrdQty" xml:"compPrdQty"`
  Status string `json:"addUseYn" bson:"addUseYn" xml:"addUseYn"`
}

type StreetProduct struct {
  SaleType string `json:"selMthdCd" bson:"selMthdCd" xml:"selMthdCd"`
  Category int `json:"dispCtgrNo" bson:"dispCtgrNo" xml:"dispCtgrNo"`
  ServiceType string `json:"prdTypeCd" bson:"prdTypeCd" xml:"prdTypeCd"`
  Name string `json:"prdNm" bson:"prdNm" xml:"prdNm"`
  Condition string `json:"prdStatCd" bson:"prdStatCd" xml:"prdStatCd"`
  Weight string `json:"prdWght" bson:"prdWght" xml:"prdWght"`

  Image01 string `json:"prdImage01" bson:"prdImage01" xml:"prdImage01"`
  Image02 string `json:"prdImage02" bson:"prdImage02" xml:"prdImage02"`
  Image03 string `json:"prdImage03" bson:"prdImage03" xml:"prdImage03"`
  Image04 string `json:"prdImage04" bson:"prdImage04" xml:"prdImage04"`
  Image05 string `json:"prdImage05" bson:"prdImage05" xml:"prdImage05"`
  Image06 string `json:"prdImage06" bson:"prdImage06" xml:"prdImage06"`
  Image07 string `json:"prdImage07" bson:"prdImage07" xml:"prdImage07"`
  Image08 string `json:"prdImage08" bson:"prdImage08" xml:"prdImage08"`

  HtmlDetail string `json:"htmlDetail" bson:"htmlDetail" xml:"htmlDetail"`
  SellTermUseYn string `json:"selTermUseYn" bson:"selTermUseYn" xml:"selTermUseYn"`
  StartDate string `json:"aplBgnDy" bson:"aplBgnDy" xml:"aplBgnDy"`
  EndDate string `json:"aplEndDy" bson:"aplEndDy" xml:"aplEndDy"`

  Price int `json:"selPrc" bson:"selPrc" xml:"selPrc"`
  ProductStock int `json:"prdSelQty" bson:"prdSelQty" xml:"prdSelQty"`
  AmountOfPurchase int `json:"selMinLimitQty" bson:"selMinLimitQty" xml:"selMinLimitQty"`
  Contact string `json:"asDetail" bson:"asDetail" xml:"asDetail"`
  DeliveryCondition string `json:"dlvCndtSeq" bson:"dlvCndtSeq" xml:"dlvCndtSeq"`
  ReturnExchange string `json:"rtngExchDetail" bson:"rtngExchDetail" xml:"rtngExchDetail"`
  VatCode string `json:"suplDtyfrPrdClfCd" bson:"suplDtyfrPrdClfCd" xml:"suplDtyfrPrdClfCd"`

  SkuList []StreetSku `json:"ProductComponent" bson:"ProductComponent" xml:"ProductComponent"`
}

type StreetProductRequest struct {
  Product StreetProduct `json:"Product" bson:"Product" xml:"Product"`
}

/// Street Config
/// ---------------------------

type StreetConfig struct {
  Url string `json:"url" bson:"url"`
  Key string `json:"key" bson:"key"`
  Version string `json:"version" bson:"version"`
}

type StreetMsgRespond struct {
  ProductNo string `json:"productNo" bson:"productNo" xml:"productNo"`
  Message string `json:"Message" bson:"Message" xml:"Message"`
  Code int `json:"resultCode" bson:"resultCode" xml:"resultCode"`
}

type StreetRespond struct {
  Product StreetMsgRespond `json:"Product" bson:"Product" xml:"Product"`
}
