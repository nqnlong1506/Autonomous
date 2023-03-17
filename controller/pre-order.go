package controller

import (
	"Autonomous/model"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePreOrder(preOrder model.PreOrder) (*model.PreOrder, error) {
	// generate new order id
	id, err := genPreOrderID()
	if err != nil {
		return nil, err
	}
	// check existing customer
	filterCustomer := bson.M{"customer_id": preOrder.CustomerID}
	existingCustomer := model.CustomerCollection.FindOne(ctx, filterCustomer)
	if existingCustomer.Err() != nil {
		return nil, fmt.Errorf("this customer does not exist")
	}
	// check existing product
	filterProduct := bson.M{"product_id": preOrder.ProductID}
	existingProduct := model.ProductCollection.FindOne(ctx, filterProduct)
	if existingProduct.Err() != nil {
		return nil, fmt.Errorf("this product does not exist")
	}
	// check pre-order products
	var product model.Product
	existingProduct.Decode(&product)
	if product.Status != model.PRE_ORDER {
		return nil, fmt.Errorf("you cannot pre-order this product")
	}
	// create new pre-order
	newPreOrder := model.PreOrder{
		PreOrderID:   id,
		Number:       preOrder.Number,
		CustomerID:   preOrder.CustomerID,
		CustomerName: preOrder.CustomerName,
		ProductID:    preOrder.ProductID,
		ProductName:  product.ProductName,
		VendorID:     product.VendorID,
	}
	_, errInsert := model.PreOrderCollection.InsertOne(ctx, newPreOrder)
	if errInsert != nil {
		return nil, errInsert
	}

	return &newPreOrder, nil
}

func CreatePreOrderForHugeQuantity(preOrder model.PreOrder) (*model.PreOrder, error) {
	// generate new order id
	id, err := genPreOrderID()
	if err != nil {
		return nil, err
	}
	// check existing customer
	filterCustomer := bson.M{"customer_id": preOrder.CustomerID}
	existingCustomer := model.CustomerCollection.FindOne(ctx, filterCustomer)
	if existingCustomer.Err() != nil {
		return nil, fmt.Errorf("this customer does not exist")
	}
	// check existing product
	filterProduct := bson.M{"product_id": preOrder.ProductID}
	existingProduct := model.ProductCollection.FindOne(ctx, filterProduct)
	if existingProduct.Err() != nil {
		return nil, fmt.Errorf("this product does not exist")
	}
	// check buy-now products
	var product model.Product
	existingProduct.Decode(&product)

	// create new order
	newPreOrder := model.PreOrder{
		PreOrderID:   id,
		Number:       preOrder.Number,
		CustomerID:   preOrder.CustomerID,
		CustomerName: preOrder.CustomerName,
		ProductID:    preOrder.ProductID,
		ProductName:  product.ProductName,
		VendorID:     product.VendorID,
	}
	_, errInsert := model.PreOrderCollection.InsertOne(ctx, newPreOrder)
	if errInsert != nil {
		return nil, errInsert
	}

	return &newPreOrder, nil
}

func genPreOrderID() (int, error) {
	opts := options.Find().SetSort(bson.M{"pre_order_id": -1})
	cursor, err := model.PreOrderCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return 0, fmt.Errorf("create new order failed")
	}
	var list []model.PreOrder
	err1 := cursor.All(ctx, &list)
	if err1 != nil {
		return 0, fmt.Errorf("create new order failed")
	}

	if len(list) <= 0 {
		return 10001, nil
	}

	return list[0].PreOrderID + 1, nil
}

func HandlePreOrder(listPreOrder []model.PreOrder) error {
	var listPreOrderID []int
	for _, p := range listPreOrder {
		err := CreateOrderFromPreOrder(p.CustomerID, p.ProductID)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
		listPreOrderID = append(listPreOrderID, p.PreOrderID)
	}

	_, errUpdate := model.PreOrderCollection.UpdateMany(ctx, bson.M{
		"pre_order_id": bson.M{
			"$in": listPreOrderID,
		},
	}, bson.M{
		"$set": bson.M{
			"is_done": true,
		},
	})
	if errUpdate != nil {
		log.Fatal(errUpdate.Error())
		return errUpdate
	}
	return nil
}

func GetAllPreOrderByVendorID(vendorID int) ([]model.PreOrder, error) {
	cursor, err := model.PreOrderCollection.Find(ctx, bson.M{
		"vendor_id": vendorID,
		"is_done":   false,
	})
	if err != nil {
		return nil, fmt.Errorf("create new order failed")
	}
	var list []model.PreOrder
	err1 := cursor.All(ctx, &list)
	if err1 != nil {
		return nil, fmt.Errorf("create new order failed")
	}
	return list, nil
}

func GetPreOrderByID(preOrderID int) (*model.PreOrder, error) {
	filter := bson.M{
		"pre_order_id": preOrderID,
		"is_done":      false,
	}
	result := model.PreOrderCollection.FindOne(ctx, filter)
	var data model.PreOrder
	result.Decode(&data)
	if result.Err() != nil {
		return nil, result.Err()
	}

	return &data, nil
}

func ProcessPreOrder(preOrderID int) error {
	filter := bson.M{
		"pre_order_id": preOrderID,
		"is_done":      false,
	}
	_, err := model.PreOrderCollection.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"is_done": true,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
