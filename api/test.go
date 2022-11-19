package api

import (
	"Autonomous/model"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx) error {
	customer := model.Customer{
		CustomerName: "nqnlong",
		Username:     "nqnlong1506",
		Password:     "nowright1506@",
	}

	response, _ := json.Marshal(customer)

	return c.SendString(string(response))
}
