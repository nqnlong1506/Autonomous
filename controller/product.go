package controller

import (
	"Autonomous/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateProduct(product model.Product) error {
	// query existing product
	existingProduct := model.ProductCollection.FindOne(ctx, bson.M{
		"$or": []bson.M{
			{"product_id": product.ProductID},
			{"product_name": product.ProductName},
			{"sku": product.Sku},
		},
	})
	if existingProduct.Err() == nil {
		return fmt.Errorf("this product already existed")
	}
	// check existing vendor
	filter := bson.M{"vendor_id": product.VendorID}
	existingVendor := model.VendorCollection.FindOne(ctx, filter)
	if existingVendor.Err() != nil {
		return fmt.Errorf("this vendor does not exist")
	}
	// insert new product
	_, err := model.ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	// update vendor
	var vendor model.Vendor
	err1 := existingVendor.Decode(&vendor)
	if err1 != nil {
		return fmt.Errorf("can not parse existing vendor")
	}
	var listProducts []string
	listProducts = append(listProducts, vendor.ListProducts...)
	listProducts = append(listProducts, product.Sku)
	_, err2 := model.VendorCollection.UpdateOne(ctx, filter, bson.M{
		"$set": bson.M{
			"list_products": listProducts,
		},
	})
	if err2 != nil {
		return err2
	}
	return nil
}

func ChangeProductToBestSeller(productID int) error {
	// query existing product
	existingProduct := model.ProductCollection.FindOne(ctx, bson.M{
		"product_id": productID,
	})
	if existingProduct.Err() != nil {
		return fmt.Errorf("this product does not exist")
	}
	// check already best seller
	var product model.Product
	err1 := existingProduct.Decode(&product)
	if err1 != nil {
		return fmt.Errorf("can not parse existing vendor")
	}
	if product.IsBestSeller && product.Minimum == 10 {
		return nil
	}
	// update product
	_, err := model.ProductCollection.UpdateOne(ctx, bson.M{
		"product_id": productID,
	}, bson.M{
		"$set": bson.M{
			"is_best_seller": true,
			"minimum":        10,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func ChangeProductToNonBestSeller(productID int) error {
	// query existing product
	existingProduct := model.ProductCollection.FindOne(ctx, bson.M{
		"product_id": productID,
	})
	if existingProduct.Err() != nil {
		return fmt.Errorf("this product does not exist")
	}
	// update product
	_, err := model.ProductCollection.UpdateOne(ctx, bson.M{
		"product_id": productID,
	}, bson.M{
		"$set": bson.M{
			"is_best_seller": false,
			"minimum":        5,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetProductInfoBySku(sku string, output *model.Product) error {
	singleResult := model.ProductCollection.FindOne(ctx, bson.M{
		"sku": sku,
	})
	if singleResult.Err() != nil {
		return fmt.Errorf("this product does not exist")
	}
	err := singleResult.Decode(output)
	if err != nil {
		return err
	}

	return nil
}

func GetAllProduct() ([]model.Product, error) {
	cursor, err := model.ProductCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var listProducts []model.Product
	err = cursor.All(ctx, &listProducts)
	if err != nil {
		return nil, err
	}

	return listProducts, nil
}

func ImportProducts(sku string, number int) error {
	// check existing product
	existingProduct := model.ProductCollection.FindOne(ctx, bson.M{
		"sku": sku,
	})
	if existingProduct.Err() != nil {
		return fmt.Errorf("this product does not exist")
	}

	var product model.Product
	err := existingProduct.Decode(&product)
	if err != nil {
		return err
	}
	// handle preorder
	opts := options.Find().SetSort(bson.M{"pre_order_id": 1}).SetLimit(int64(number + product.Available - 1))
	cursor, errFind := model.PreOrderCollection.Find(ctx, bson.M{
		"product_id": product.ProductID,
		"is_done":    false,
	}, opts)
	if errFind != nil {
		return errFind
	}
	var listPreOrder []model.PreOrder
	cursor.All(ctx, &listPreOrder)

	// var listPreOrderID []int
	// for _, p := range listPreOrder {
	// 	listPreOrderID = append(listPreOrderID, p.PreOrderID)
	// }

	go HandlePreOrder(listPreOrder)

	if len(listPreOrder) == number {
		return nil
	}

	// update product available
	_, errUpdate := model.ProductCollection.UpdateOne(ctx, product, bson.M{
		"$inc": bson.M{
			"available": number - len(listPreOrder),
		},
		"$set": bson.M{
			"status": model.BUY_NOW,
		},
	})
	if errUpdate != nil {
		return errUpdate
	}

	// set received_mail to false
	if product.Available+number > product.Minimum {
		filter := bson.M{
			"vendor_id": product.VendorID,
		}
		updater := bson.M{
			"$set": bson.M{
				"received_mail": false,
			},
		}
		model.VendorCollection.UpdateOne(ctx, filter, updater)
	}

	return nil
}

func UpdateProductImageLink(product model.Product) error {
	_, err := model.ProductCollection.UpdateOne(ctx, bson.M{
		"product_id": product.ProductID,
	}, bson.M{
		"$set": bson.M{
			"image_link": product.ImageLink,
		},
	})

	if err != nil {
		return err
	}

	return nil
}
