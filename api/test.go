package api

import (
	"Autonomous/model"
	"context"
	"fmt"
	"log"

	mongoClient "Autonomous/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

func Test(c *fiber.Ctx) error {
	customer := model.Customer{
		CustomerName: "nqnlong",
		Username:     "nqnlong1506",
		Password:     "nowright1506@",
	}

	collection := mongoClient.Database.Collection("tasks")
	// _, err := collection.InsertOne(ctx, customer)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	result := collection.FindOne(ctx, customer)
	var data model.Customer
	result.Decode(&data)
	fmt.Println(data)

	response := model.Response{
		Status:   "OK",
		Data:     customer,
		Messsage: "Test Successfully",
	}
	return c.Status(200).JSON(response)
}

func SendEmail(c *fiber.Ctx) error {
	model.VendorCollection.UpdateMany(ctx, bson.M{}, bson.M{
		"$set": bson.M{
			"received_mail": false,
		},
	})
	return c.SendString("Sending Email Successfully!")
}

func InsertCustomer(c *fiber.Ctx) error {
	result, err := model.ProductCollection.UpdateMany(ctx, bson.M{}, bson.M{
		"$set": bson.M{
			"available": 12,
			"status":    model.BUY_NOW,
		},
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(result)
	return nil
}

func InsertProduct(c *fiber.Ctx) error {
	result, err := model.ProductCollection.InsertOne(ctx, model.Product{})

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(result)
	return nil
}

func SetPrice(c *fiber.Ctx) error {
	result, err := model.ProductCollection.UpdateMany(ctx, bson.M{}, bson.M{
		"$set": bson.M{
			"price": 99,
		},
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(result)
	return c.Status(200).JSON(model.Response{
		Status: "OK",
	})
}
