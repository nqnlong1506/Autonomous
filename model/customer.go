package model

type Customer struct {
	CustomerName string `json:"customerName" bson:"customer_name,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
	Password     string `jsonn:"password" bson:"password,omitempty"`
}
