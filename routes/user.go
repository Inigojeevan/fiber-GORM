package routes

import (
	"strconv"

	"github.com/Inigojeevan/fiber-GORM/database"
	"github.com/Inigojeevan/fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	if err != nil {
		return err
	}

	user := models.User{
		Id:        id,
		FirstName: data["firstName"],
		LastName:  data["lastName"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
