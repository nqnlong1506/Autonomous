package controller

import (
	"Autonomous/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateOrder(order model.Order) (*model.Order, error) {
	// generate new order id
	id, err := genOrderID()
	if err != nil {
		return nil, err
	}
	// check existing customer
	filterCustomer := bson.M{"customer_id": order.CustomerID}
	existingCustomer := model.CustomerCollection.FindOne(ctx, filterCustomer)
	if existingCustomer.Err() != nil {
		return nil, fmt.Errorf("this customer does not exist")
	}
	// check existing product
	filterProduct := bson.M{"product_id": order.ProductID}
	existingProduct := model.ProductCollection.FindOne(ctx, filterProduct)
	if existingProduct.Err() != nil {
		return nil, fmt.Errorf("this product does not exist")
	}
	// check buy-now products
	var product model.Product
	existingProduct.Decode(&product)
	if product.Status != model.BUY_NOW {
		return nil, fmt.Errorf("you cannot buy-now this product")
	}
	// create new order
	newOrder := model.Order{
		OrderID:    id,
		CustomerID: order.CustomerID,
		ProductID:  order.ProductID,
	}
	_, errInsert := model.OrderCollection.InsertOne(ctx, newOrder)
	if errInsert != nil {
		return nil, errInsert
	}
	// update product
	go updateAvailableAfterBuy(product)
	return &newOrder, nil
}

func CreateOrderFromPreOrder(customerID int, productID int) error {
	// generate new order id
	id, err := genOrderID()
	if err != nil {
		return err
	}
	// create new order
	newOrder := model.Order{
		OrderID:    id,
		CustomerID: customerID,
		ProductID:  productID,
	}
	_, errInsert := model.OrderCollection.InsertOne(ctx, newOrder)
	if errInsert != nil {
		return errInsert
	}
	return nil
}

func genOrderID() (int, error) {
	opts := options.Find().SetSort(bson.M{"order_id": -1})
	cursor, err := model.OrderCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return 0, fmt.Errorf("create new order failed")
	}
	var list []model.Order
	err1 := cursor.All(ctx, &list)
	if err1 != nil {
		return 0, fmt.Errorf("create new order failed")
	}

	if len(list) <= 0 {
		return 10001, nil
	}

	return list[0].OrderID + 1, nil
}

func updateAvailableAfterBuy(product model.Product) error {
	updater := bson.M{
		"$inc": bson.M{
			"available": -1,
		},
	}
	if product.Available <= 2 {
		updater["$set"] = bson.M{
			"status": model.PRE_ORDER,
		}
	}
	_, errUpdate := model.ProductCollection.UpdateOne(ctx, product, updater)
	if errUpdate != nil {
		fmt.Println(errUpdate.Error())
	}

	// send email to vendor
	if product.Available-1 == product.Minimum {
		SendMailForUpdatingInventory(product.VendorID, product.ProductName, product.Sku)
	}
	return nil
}
