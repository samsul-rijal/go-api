package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samsul-rijal/go-api/controller"
	"github.com/samsul-rijal/go-api/middleware"
)

func RouteInit(r *fiber.App) {
	r.Static("/public/assets", "/ /assets")
	r.Get("/user", middleware.Auth, controller.UserControllerGetAll)
	r.Get("/user/:id", controller.UserControllerGetById)
	r.Post("/user", controller.UserControllerCreate)
	r.Put("/user/:id", controller.UserControllerUpdate)
	r.Delete("/user/:id", controller.UserControllerDelete)

	r.Post("/login", controller.LoginController)
}