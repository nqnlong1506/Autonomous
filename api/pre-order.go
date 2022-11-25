package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func CreatePreOrder(c *fiber.Ctx) error {
	var preOrder model.PreOrder
	if err := c.BodyParser(&preOrder); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	if preOrder.CustomerID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer id needed",
		})
	}
	if preOrder.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	newOrder, err := controller.CreatePreOrder(preOrder)
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
