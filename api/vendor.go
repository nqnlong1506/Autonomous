package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func GetVendorInfo(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "vendor id needed",
		})
	}
	return c.SendString(id)
}

func CreateVendorInfo(c *fiber.Ctx) error {
	var vendor model.Vendor
	if err := c.BodyParser(&vendor); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if vendor.VendorID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "vendor id needed",
		})
	}

	if vendor.VendorName == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "vendor name needed",
		})
	}

	if vendor.Username == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "username needed",
		})
	}

	if vendor.Password == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "password needed",
		})
	}

	if vendor.Email == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "email needed",
		})
	}

	err := controller.CreateVendorInfo(vendor)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     vendor,
		Messsage: "Create vendor successfully",
	})
}
