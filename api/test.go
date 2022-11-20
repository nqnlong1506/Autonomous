package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	customer := model.Customer{
		CustomerName: "nqnlong",
		Username:     "nqnlong1506",
		Password:     "nowright1506@",
	}

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
