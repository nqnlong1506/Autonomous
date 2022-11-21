package main

import (
	"Autonomous/api"
	mongoClient "Autonomous/mongo"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection
var ctx = context.TODO()

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

	// connect mongodb
	{
		uri := "mongodb://localhost:27017/"
		database := "autonomous"
		err := mongoClient.InitializeClient(ctx, uri, database)
		if err != nil {
			log.Fatal(err)
		}
	}

	app.Get("/", infoHandle)

	{
		app.Get("/test", api.Test)
		app.Get("/send_email", api.SendEmail)
	}

	app.Listen(":3000")
}
