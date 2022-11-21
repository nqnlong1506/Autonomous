package api

import (
	"Autonomous/controller"
	"Autonomous/model"
	"context"
	"fmt"

	mongoClient "Autonomous/mongo"

	"github.com/gofiber/fiber/v2"
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
	err := controller.SendMailForUpdatingInventory("nqnl1506@gmail.com", "nha sach sai gon", "sach tieng anh")
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.SendString("Sending Email Successfully!")
}
