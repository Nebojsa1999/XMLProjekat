package controller

import (
	"context"
	"github.com/Nebojsa1999/XMLProjekat/backend/user_service/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"time"
)

const SecretKey = "tajna-tima-14"

func RegisterANewUser(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if isUsernameAlreadyTaken(data["username"]) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "username is already taken, please choose another one",
		})
	}

	if isEmailAlreadyTaken(data["email"]) {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "email is already taken, please choose another one",
		})
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	dateOfBirth, _ := time.Parse("2010-01-30", data["date-of-birth"])

	user := model.User{
		Id:             0,
		Username:       data["username"],
		Password:       password,
		FirstName:      data["first_name"],
		LastName:       data["last_name"],
		Email:          data["email"],
		Gender:         model.Gender(data["gender"]),
		DateOfBirth:    dateOfBirth,
		Biography:      data["biography"],
		WorkExperience: data["work-experience"],
		Education:      data["education"],
		Skills:         data["skills"],
		Interests:      data["interests"],
	}

	usersCollection := GetAllUsers()
	if _, err := usersCollection.InsertOne(context.TODO(), user); err != nil {
		log.Fatal(err)
	}

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	existingUser := GetUserByUsername(data["username"])
	if existingUser.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "existing user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(existingUser.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(existingUser.Id)),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "successfully logged in",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "successfully logged out",
	})
}
