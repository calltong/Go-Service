package router

import (
  "gopkg.in/mgo.v2/bson"
)

type Ordered struct {
  Image string `json:"image" bson:"image"`
  ProductId bson.ObjectId `json:"product_id" bson:"product_id"`
  ColorId bson.ObjectId `json:"color_id" bson:"color_id"`
	SizeId bson.ObjectId `json:"size_id" bson:"size_id"`
  Name string `json:"name" bson:"name"`
  Color string `json:"color" bson:"color"`
  Size string `json:"size" bson:"size"`
  Price int `json:"price" bson:"price"`
  Quantity int `json:"quantity" bson:"quantity"`
}

type PaySlip struct {
	Slip string `json:"slip" bson:"slip"`
  Updated bool `json:"updated" bson:"updated"`
}

type PaymentOrder struct {
	Type string `json:"type" bson:"type"`
	Data PaySlip `json:"data" bson:"data"`
}

type AddressOrdered struct {
  Name string `json:"name" bson:"name"`
  Address string `json:"address" bson:"address"`
	City string `json:"city" bson:"city"`
	Postcode string `json:"postcode" bson:"postcode"`
	Mobile string `json:"mobile" bson:"mobile"`
  Email string `json:"email" bson:"email"`
}

type StatusUpdated struct {
  Status string `json:"status" bson:"status"`
	UpdatedAt int64 `json:"updated_at" bson:"updated_at"`
}

type SummaryOrder struct {
	Discount int `json:"discount" bson:"discount"`
	Shipping int `json:"shipping" bson:"shipping"`
  Total int `json:"total" bson:"total"`
}

type Order struct {
	Id bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	CustomerId bson.ObjectId `json:"customer_id" bson:"customer_id,omitempty"`
  PromotionId bson.ObjectId `json:"promotion_id" bson:"promotion_id,omitempty"`
	OrderList []Ordered `json:"list" bson:"list"`
  Payment PaymentOrder `json:"payment" bson:"payment"`
  Summary SummaryOrder `json:"summary" bson:"summary"`
  Address AddressOrdered `json:"shipping" bson:"shipping"`
  StatusList []StatusUpdated `json:"status_list" bson:"status_list"`
  Status string `json:"status" bson:"status"`
  TrackingCode string `json:"tracking_code" bson:"tracking_code"`
  CreatedAt int64 `json:"created_at" bson:"created_at"`
}
