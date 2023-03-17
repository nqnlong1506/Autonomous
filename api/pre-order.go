package api

import (
	"Autonomous/controller"
	"Autonomous/model"
	"strconv"

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

	// check customer id
	if preOrder.CustomerID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer id needed",
		})
	}
	// check customer name
	if preOrder.CustomerName == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer name needed",
		})
	}
	// check product id
	if preOrder.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}
	// check quantity
	if preOrder.Number <= 0 {
		preOrder.Number = 1
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

func ListAllPreOrder(c *fiber.Ctx) error {
	vendorID := c.Query("id")
	if vendorID == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Ivalid",
			Messsage: "vendor id needed",
		})
	}

	id, _ := strconv.Atoi(vendorID)

	var list []model.PreOrder
	list, err := controller.GetAllPreOrderByVendorID(id)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     list,
		Messsage: "get list of pre-order successfully",
	})
}

func GetPreOrderByID(c *fiber.Ctx) error {
	preOrderId := c.Query("id")
	if preOrderId == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Ivalid",
			Messsage: "vendor id needed",
		})
	}

	id, _ := strconv.Atoi(preOrderId)

	preOrder, err := controller.GetPreOrderByID(id)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     preOrder,
		Messsage: "get list of pre-order successfully",
	})
}

func ProcessPreOrder(c *fiber.Ctx) error {
	var preOrder struct {
		PreOrderID int `json:"preOrderID"`
	}
	if err := c.BodyParser(&preOrder); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if preOrder.PreOrderID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "Pre Order ID needed",
		})
	}

	err := controller.ProcessPreOrder(preOrder.PreOrderID)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "process pre-order successfully",
	})
}
