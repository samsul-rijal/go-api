package controller

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samsul-rijal/go-api/config/database"
	"github.com/samsul-rijal/go-api/model/entity"
	"github.com/samsul-rijal/go-api/model/request"
	"github.com/samsul-rijal/go-api/utils"
)

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)	
}

func UserControllerCreate(ctx *fiber.Ctx) error {
	user := new(entity.User)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message":"failed",
			"error": errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Debug().Create(&newUser).Error

	if errCreateUser != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message":"Failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message":"success",
		"data": newUser,
	})
}

func UserControllerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message":"user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID: user.ID,
	// 	Name: user.Name,
	// 	Email: user.Email,
	// 	Address: user.Address,
	// 	Phone: user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.Status(404).JSON(fiber.Map{
		"message":"success",
		"data": user,
	})
}

func UserControllerUpdate(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message":"user not found",
		})
	}

	userRequest := new(request.UserUpdateRequest)	
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message":"bad request",
		})
	}

	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}

	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil{
		return ctx.Status(500).JSON(fiber.Map{
			"message":"internal server error",
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"message":"success",
		"data": user,
	})
}

func UserControllerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": fmt.Sprintf("user %s was deleted",userId),
		})
	}

	errDelete := database.DB.Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message":"internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": fmt.Sprintf("user %s was deleted",userId),
	})
}