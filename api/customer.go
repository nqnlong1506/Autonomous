package api

import (
	"Autonomous/controller"
	"Autonomous/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CreateCustomer(c *fiber.Ctx) error {
	var customer model.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "Invalid",
		})
	}

	if customer.CustomerName == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "customer name needed",
		})
	}
	if customer.Username == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "username needed",
		})
	}
	if customer.Password == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "password needed",
		})
	}

	fmt.Println(customer)
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "OK",
	})
}

func LoginCustomer(c *fiber.Ctx) error {
	var loginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&loginForm); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if loginForm.Username == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "username needed",
		})
	}

	if loginForm.Password == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "password needed",
		})
	}
	fmt.Println("login:", loginForm)

	customer, err := controller.LoginCustomer(loginForm.Username, loginForm.Password)
	if err != nil {
		return c.Status(203).JSON(model.Response{
			Status:   "Incorrect",
			Messsage: "incorrect username or password",
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "OK",
		Data:     customer,
	})
}
