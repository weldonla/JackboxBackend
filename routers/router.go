package router

import (
	"github.com/weldonla/FourLeafPortalApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/login", controllers.Login)
	app.Post("/api/register", controllers.Register)
	app.Get("/api/users", controllers.GetUserList)
}
