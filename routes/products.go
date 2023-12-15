package routes

import (
	"errors"

	"github.com/Inigojeevan/fiber-GORM/database"
	"github.com/Inigojeevan/fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(product models.Product) Product {
	return Product{
		ID:           product.ID,
		Name:         product.Name,
		SerialNumber: product.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetAllProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)

	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("User doesnt exists")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Ensure to put the id as int")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Ensure to enter the id as int")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdatedProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var update UpdatedProduct

	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = update.Name
	product.SerialNumber = update.SerialNumber

	database.Database.Db.Save(&product)

	resonseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(resonseProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Ensure to check the id as int")
	}
	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully deleted the product")
}
