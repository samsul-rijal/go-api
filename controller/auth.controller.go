package controller

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samsul-rijal/go-api/config/database"
	"github.com/samsul-rijal/go-api/model/entity"
	"github.com/samsul-rijal/go-api/model/request"
	"github.com/samsul-rijal/go-api/utils"
)

func LoginController(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message":"failed",
			"error": errValidate.Error(),
		})
	}

	// Check email
	var user entity.User
	err := database.DB.Debug().First(&user, "email = ?", loginRequest.Email).Error

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"wrong email or password!",
		})
	}

	// Check password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"wrong email or password",
		})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix() // 2 menit expired

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil{
		log.Println(errGenerateToken)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthorized",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})

}