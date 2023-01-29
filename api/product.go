package api

import (
	"Autonomous/controller"
	"Autonomous/model"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if product.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}
	if product.ProductName == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product name needed",
		})
	}
	if product.VendorID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "vendor id needed",
		})
	}
	if product.Sku == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product sku needed",
		})
	}
	if product.Price <= 0 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product price needed",
		})
	}
	if product.Available <= 1 {
		product.Available = 1
	}
	if product.IsBestSeller {
		product.Minimum = 10
	} else {
		product.Minimum = 5
	}
	if product.Available <= 1 {
		product.Status = model.PRE_ORDER
	} else {
		product.Status = model.BUY_NOW
	}

	err := controller.CreateProduct(product)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "create product successfully",
	})
}

func UpdateBestSellerProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if product.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	err := controller.ChangeProductToBestSeller(product.ProductID)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "update product successfully",
	})
}

func UpdateNonBestSellerProduct(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if product.ProductID <= 1 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	err := controller.ChangeProductToNonBestSeller(product.ProductID)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "update product successfully",
	})
}

func GetProductInfo(c *fiber.Ctx) error {
	sku := c.Query("sku")
	if sku == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Ivalid",
			Messsage: "sku product needed",
		})
	}

	var product model.Product
	err := controller.GetProductInfoBySku(sku, &product)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     product,
		Messsage: "update product successfully",
	})
}

func GetAllProduct(c *fiber.Ctx) error {

	listProducts, err := controller.GetAllProduct()
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Data:     listProducts,
		Messsage: "get all product successfully",
	})
}

func ImportProducts(c *fiber.Ctx) error {
	var products struct {
		Sku    string `json:"sku"`
		Number int    `json:"number"`
	}
	if err := c.BodyParser(&products); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if products.Sku == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "sku product needed",
		})
	}
	if products.Number <= 0 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "number of products needed",
		})
	}

	err := controller.ImportProducts(products.Sku, products.Number)
	if err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}
	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "import products successfully",
	})
}

func InsertImage(c *fiber.Ctx) error {
	var product model.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: err.Error(),
		})
	}

	if product.ProductID <= 0 {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product id needed",
		})
	}

	if product.ImageLink == "" {
		return c.Status(400).JSON(model.Response{
			Status:   "Invalid",
			Messsage: "product image link needed",
		})
	}

	errUpdate := controller.UpdateProductImageLink(product)
	if errUpdate != nil {
		return c.Status(400).JSON(model.Response{
			Status:   "Error",
			Messsage: errUpdate.Error(),
		})
	}

	return c.Status(200).JSON(model.Response{
		Status:   "OK",
		Messsage: "update image link successfully",
	})
}
