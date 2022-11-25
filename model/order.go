package model

import "go.mongodb.org/mongo-driver/mongo"

type Order struct {
	OrderID    int `json:"orderID,omitempty" bson:"order_id,omitempty"`
	CustomerID int `json:"customerID,omitempty" bson:"customer_id,omitempty"`
	ProductID  int `json:"productID,omitempty" bson:"product_id,omitempty"`
}

var OrderCollection *mongo.Collection

func InitOrderCollection(database mongo.Database) {
	OrderCollection = database.Collection("order")
}
