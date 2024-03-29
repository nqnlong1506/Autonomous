package controller

import (
	"Autonomous/model"
	"fmt"
	"net/smtp"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	companyEmail = "ngocdang1607@gmail.com"
	appPassword  = "ucljeijwbchnboqy"
)

func SendMailForUpdatingInventory(vendorID int, product model.Product) error {
	// check existing vendor
	existingVendor := model.VendorCollection.FindOne(ctx, bson.M{
		"vendor_id": vendorID,
	})
	if existingVendor.Err() != nil {
		return fmt.Errorf("this vendor already existed")
	}
	var vendor model.Vendor
	err := existingVendor.Decode(&vendor)
	if err != nil {
		return err
	}

	// send mail
	to := []string{vendor.Email}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: Update inventory urgently for more orders!!!\n"
	body := fmt.Sprintf("Hi %s team,\n\nYour product %s (%s) is selling fastly and going to be out of stock. Please help us to update the stock as soon as possible if your inventory is still available to sell. If not, when the product reach the lowest stock, we will make the one become Pre-Order for more 1 months and you can track the order later.\nLooking forward to hearing from you soon.\n\nThanks and best regards,\nGalvin.", vendor.VendorName, product.ProductName, product.Sku)
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", companyEmail, appPassword, host)

	errSendMail := smtp.SendMail(address, auth, companyEmail, to, message)
	if errSendMail != nil {
		return errSendMail
	}
	// update product
	updater := bson.M{
		"$set": bson.M{
			"received_mail": true,
		},
	}
	_, errUpdate := model.ProductCollection.UpdateMany(ctx, bson.M{
		"product_id": product.ProductID,
	}, updater)
	if errUpdate != nil {
		fmt.Print(errUpdate.Error())
		return errUpdate
	}
	return nil
}
