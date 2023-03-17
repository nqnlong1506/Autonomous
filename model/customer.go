package model

import "go.mongodb.org/mongo-driver/mongo"

type Customer struct {
	CustomerID   int    `json:"customerID" bson:"customer_id"`
	CustomerName string `json:"customerName" bson:"customer_name"`
	Username     string `json:"username" bson:"username"`
	Password     string `jsonn:"password" bson:"password"`
}

var CustomerCollection *mongo.Collection

func InitCustomerCollection(database mongo.Database) {
	CustomerCollection = database.Collection("customer")
}
