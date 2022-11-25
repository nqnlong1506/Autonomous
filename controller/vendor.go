package controller

import (
	"Autonomous/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

func CreateVendorInfo(vendor model.Vendor) error {
	// query existing vendor
	existingVendor := model.VendorCollection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"vendor_id": vendor.VendorID},
			{"vendor_name": vendor.VendorName},
			{"username": vendor.Username},
		},
	})
	if existingVendor.Err() == nil {
		return fmt.Errorf("this vendor already existed")
	}
	_, err := model.VendorCollection.InsertOne(ctx, vendor)
	if err != nil {
		return err
	}

	return nil
}
