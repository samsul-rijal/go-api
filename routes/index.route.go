package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samsul-rijal/go-api/controller"
)

func RouteInit(r *fiber.App) {
	r.Static("/public/assets", "/public/assets")
	r.Get("/user", controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Delete("/user/:id", controller.UserControllerDelete)
}