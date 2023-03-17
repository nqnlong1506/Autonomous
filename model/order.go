package model

import "go.mongodb.org/mongo-driver/mongo"

type Order struct {
	OrderID      int    `json:"orderID,omitempty" bson:"order_id,omitempty"`
	Number       int    `json:"number,omitempty" bson:"number,omitempty"`
	CustomerID   int    `json:"customerID,omitempty" bson:"customer_id,omitempty"`
	CustomerName string `json:"customerName,omitempty" bson:"customer_name,omitempty"`
	ProductID    int    `json:"productID,omitempty" bson:"product_id,omitempty"`
	ProductName  string `json:"productName,omitempty" bson:"product_name,omitempty"`
}

var OrderCollection *mongo.Collection

func InitOrderCollection(database mongo.Database) {
	OrderCollection = database.Collection("order")
}
