package model

import "go.mongodb.org/mongo-driver/mongo"

type PreOrder struct {
	PreOrderID   int    `json:"preOrderID,omitempty" bson:"pre_order_id,omitempty"`
	Number       int    `json:"number,omitempty" bson:"number,omitempty"`
	CustomerID   int    `json:"customerID,omitempty" bson:"customer_id,omitempty"`
	CustomerName string `json:"customerName,omitempty" bson:"customer_name,omitempty"`
	ProductID    int    `json:"productID,omitempty" bson:"product_id,omitempty"`
	ProductName  string `json:"productName,omitempty" bson:"product_name,omitempty"`
	VendorID     int    `json:"vendorID,omitempty" bson:"vendor_id,omitempty"`
	IsDone       bool   `json:"isDone,omitempty" bson:"is_done"`
}

var PreOrderCollection *mongo.Collection

func InitPreOrderCollection(database mongo.Database) {
	PreOrderCollection = database.Collection("pre_order")
}
