package controller

import (
	"Autonomous/model"
	"fmt"

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
	// check buy-now products
	var product model.Product
	existingProduct.Decode(&product)
	if product.Status != model.PRE_ORDER {
		return nil, fmt.Errorf("you cannot pre-order this product")
	}
	// create new order
	newPreOrder := model.PreOrder{
		PreOrderID: id,
		CustomerID: preOrder.CustomerID,
		ProductID:  preOrder.ProductID,
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
	fmt.Println("handle preorder:", listPreOrder)

	var listPreOrderID []int
	for _, p := range listPreOrder {
		err := CreateOrderFromPreOrder(p.CustomerID, p.ProductID)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		listPreOrderID = append(listPreOrderID, p.PreOrderID)
	}

	fmt.Println(listPreOrderID)
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
		fmt.Println(errUpdate.Error())
		return errUpdate
	}
	return nil
}
