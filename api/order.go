package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	var order model.Order

	// get request body from front end
	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	// check customer id
	if order.CustomerID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer id needed",
		})
	}

	// check customer name
	if order.CustomerName == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer name needed",
		})
	}

	// check product id
	if order.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	// check quantity
	if order.Number <= 0 {
		order.Number = 1
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
