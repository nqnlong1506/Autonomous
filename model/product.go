package model

import "go.mongodb.org/mongo-driver/mongo"

const (
	PRE_ORDER = "Pre-Order"
	BUY_NOW   = "Buy-Now"
)

type Product struct {
	ProductID    int    `json:"productID,omitempty" bson:"product_id,omitempty"`
	ProductName  string `json:"productName,omitempty" bson:"product_name,omitempty"`
	VendorID     int    `json:"vendorID,omitempty" bson:"vendor_id,omitempty"`
	Price        int    `json:"price" bson:"price"`
	Available    int    `json:"available,omitempty" bson:"available,omitempty"`
	IsBestSeller bool   `json:"isBestSeller,omitempty" bson:"is_best_seller"`
	Minimum      int    `json:"-" bson:"minimum,omitempty"`
	Status       string `json:"status" bson:"status,omitempty"`
	Sku          string `json:"sku,omitempty" bson:"sku,omitempty"`
	ImageLink    string `json:"imageLink,omitempty" bson:"image_link"`
	ReceivedMail bool   `bson:"received_mail"`
}

var ProductCollection *mongo.Collection

func InitProductCollection(database mongo.Database) {
	ProductCollection = database.Collection("product")
}
