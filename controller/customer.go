package controller

import (
	"Autonomous/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func LoginCustomer(username, password string) (*model.Customer, error) {
	filterCustomer := bson.M{
		"username": username,
		"password": password,
	}
	result := model.CustomerCollection.FindOne(ctx, filterCustomer)
	var data model.Customer
	result.Decode(&data)
	fmt.Println(data)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &data, nil
}
