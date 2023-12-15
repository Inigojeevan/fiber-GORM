package routes

import (
	"errors"

	"github.com/Inigojeevan/fiber-GORM/database"
	"github.com/Inigojeevan/fiber-GORM/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	//this type can be used to display the necessary info
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}    //used for only needed things(slice)
	for _, user := range users { //new variable user
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func helperFunc(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user doesnt exists")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	var user models.User
	id, err := c.ParamsInt("id") //if nothing is found, it returns 0

	if err != nil {
		return c.Status(400).JSON("Ensure to check the id as int")
	}
	if err := helperFunc(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Ensure to check the id as int")
	}
	if err := helperFunc(id, &user); err != nil {
		return c.Status(400).JSON(err.Error)
	}

	type UpdateUsers struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateusers UpdateUsers

	if err := c.BodyParser(&updateusers); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateusers.FirstName
	user.LastName = updateusers.LastName

	database.Database.Db.Save(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(400).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Ensure to check the id as int")
	}
	if err := helperFunc(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(400).SendString("Successfully deleted the user")
}
