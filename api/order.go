package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	if order.CustomerID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer id needed",
		})
	}
	if order.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	newOrder, err := controller.CreateOrder(order)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     newOrder,
		Messsage: "create order succesfully",
	})
}
