package routes

import (
	"errors"
	"time"

	"github.com/Inigojeevan/fiber-GORM/database"
	"github.com/Inigojeevan/fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`    //from router
	Product   Product   `json:"product"` //from router
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,    //from User routes
		Product:   product, //from Product routes
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := helperFunc(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var product models.Product

	if err := findProduct(order.ProductRefer, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}

func findOrders(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("orders doesnt exists")
	}
	return nil
}

func GetAllOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Find(&orders)

	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}
	return c.Status(200).JSON(responseOrders)
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order

	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Ensure to check the id as int")
	}
	if err := findOrders(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	var product models.Product
	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)

	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)
	return c.Status(200).JSON(responseOrder)
}
