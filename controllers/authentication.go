package controllers

import (
	"fmt"
	"time"

	database "github.com/weldonla/FourLeafPortalApi/databases"
	"github.com/weldonla/FourLeafPortalApi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	if err := database.DB.Where("user_name = ?", data["userName"]).First(&user).Error; err != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	})

	token, err := claims.SignedString([]byte("HailAllYeRounders"))

	if user.Id.String() == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token":     token,
		"id":        user.Id,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"userName":  user.UserName,
		"email":     user.Email,
		"phone":     user.Phone,
		"password":  user.Password,
		"isAdmin":   user.IsAdmin,
	})
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var hashedPassword, hashErr = bcrypt.GenerateFromPassword([]byte(data["password"]), 10)

	isAdmin := false

	if data["isAdmin"] == "true" {
		isAdmin = true
	}

	var newUser = models.User{
		FirstName: data["firstName"],
		LastName:  data["lastName"],
		Email:     data["email"],
		Phone:     data["phone"],
		Password:  hashedPassword,
		UserName:  data["userName"],
		IsAdmin:   isAdmin,
	}

	fmt.Println("userName: " + newUser.UserName)
	fmt.Println("dataUserName: " + data["userName"])

	var user models.User
	database.DB.Where("user_name = ?", data["userName"]).First(&user)

	if user.UserName != "" {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "user already exists",
		})
	}

	database.DB.Where("email = ?", data["email"]).First(&user)
	if user.Email != "" {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "user already exists",
		})
	}

	if database.DB.Where("user_name = ?", data["userName"]).Updates(&newUser).RowsAffected == 0 {
		database.DB.Create(&newUser)
	}

	if hashErr != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not create user",
		})
	}

	return c.JSON(fiber.Map{
		"id":        user.Id,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"userName":  user.UserName,
		"email":     user.Email,
		"phone":     user.Phone,
		"password":  user.Password,
		"isAdmin":   user.IsAdmin,
	})
}

func GetUserList(c *fiber.Ctx) error {
	// var data map[string]string

	// if err := c.BodyParser(&data); err != nil {
	// 	return err
	// }

	// isAdmin := false

	// if data["isAdmin"] == "true" {
	// 	isAdmin = true
	// }

	// var inputUser = models.User{
	// 	FirstName: data["firstName"],
	// 	LastName:  data["lastName"],
	// 	Email:     data["email"],
	// 	Phone:     data["phone"],
	// 	UserName:  data["userName"],
	// 	IsAdmin:   isAdmin,
	// }

	// fmt.Println("userName: " + inputUser.UserName)
	// fmt.Println("dataUserName: " + data["userName"])

	// var user models.User
	// database.DB.Where("user_name = ?", inputUser.UserName).First(&user)

	// if user.IsAdmin != true {
	// 	c.Status(fiber.StatusBadRequest)
	// 	return c.JSON(fiber.Map{
	// 		"message": "there is something wrong with your credentials",
	// 	})
	// }
	var userList []models.User
	database.DB.Find(&userList)

	return c.JSON(fiber.Map{
		"users": userList,
	})
}
