package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to database")
	}
	DB.AutoMigrate(&User{})
}

func GetUsersList(c *fiber.Ctx) error {
	var users []User
	DB.Find(&users)

	return c.JSON(&users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	DB.Find(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(&user)
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Create(&user)

	return c.JSON(&user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User not found")
	}

	DB.Delete(&user)

	return c.SendString("User with id " + user.Email + " deleted successfully")
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	DB.First(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User not found")
	}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	DB.Save(&user)
	return c.JSON(&user)
}
