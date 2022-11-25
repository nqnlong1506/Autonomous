package main

import (
	"Autonomous/api"
	"Autonomous/model"
	mongoClient "Autonomous/mongo"
	"context"
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

	return c.Status(200).JSON(model.Response{
		Status: "OK",
		Data:   info,
	})
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
	// init collection
	{
		model.InitCustomerCollection(*mongoClient.Database)
		model.InitProductCollection(*mongoClient.Database)
		model.InitVendorCollection(*mongoClient.Database)
		model.InitOrderCollection(*mongoClient.Database)
		model.InitPreOrderCollection(*mongoClient.Database)
	}

	app.Get("/", infoHandle)
	{
		app.Get("/test", api.Test)
		app.Get("/send_email", api.SendEmail)
		app.Post("/customer/test", api.InsertCustomer)
		app.Post("/product/test", api.InsertProduct)
	}
	// vendor
	{
		app.Get("/vendor", api.GetVendorInfo)
		app.Post("/vendor/create", api.CreateVendorInfo)
	}
	// product
	{
		app.Get("/product", api.GetProductInfo)
		app.Get("/product/all", api.GetAllProduct)
		app.Post("/product/create", api.CreateProduct)
		app.Put("/product/import", api.ImportProducts)
		app.Put("/product/update/best-seller", api.UpdateBestSellerProduct)
		app.Put("/product/update/non-best-seller", api.UpdateNonBestSellerProduct)
	}
	// order
	{
		app.Post("/order/create", api.CreateOrder)
	}
	// pre-order
	{
		app.Post("/pre-order/create", api.CreatePreOrder)
	}
	// customer
	{
		// app.Post("/customer/order", api.GetOrder)
	}

	app.Listen(":3000")
}
