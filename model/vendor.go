package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Vendor struct {
	VendorID   int    `json:"vendorID,omitempty" bson:"vendor_id,omitempty"`
	VendorName string `json:"vendorName,omitempty" bson:"vendor_name,omitempty"`
	Email      string `json:"email,omitempty" bson:"email,omitempty"`
	// Address      string   `json:"address,omitempty" bson:"address,omitempty"`
	// CeoName      string   `json:"ceoName,omitempty" bson:"ceo_name"`
	ListProducts []string `json:"listProducts,omitempty" bson:"list_products"`

	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

var VendorCollection *mongo.Collection

func InitVendorCollection(database mongo.Database) {
	VendorCollection = database.Collection("vendor")
}
