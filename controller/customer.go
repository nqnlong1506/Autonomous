package controller

import (
	"Autonomous/model"

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
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &data, nil
}
