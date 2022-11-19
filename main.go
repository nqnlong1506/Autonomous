package main

import (
	"Autonomous/api"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

type apiInfo struct {
	Service  string    `json:"service"`
	Startime time.Time `json:"startTime"`
	Version  string    `json:"version"`
}

func infoHandle(c *fiber.Ctx) error {
	info := apiInfo{
		Service:  "E-commercial",
		Startime: time.Date(2022, 10, 19, 0, 0, 0, 0, time.UTC),
		Version:  "v.0.1.0",
	}

	response, _ := json.Marshal(info)

	return c.SendString(string(response))
}

func main() {
	app := fiber.New()

	app.Get("/", infoHandle)

	{
		app.Get("/test", api.Test)
	}

	app.Listen(":3000")
}
