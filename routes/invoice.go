package routes

import (
	"invoice-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupInvoceRoutes(app *fiber.App) {
	app.Get("/invoice", controllers.ListInvoices)
	app.Post("/invoice", controllers.CreateInvoice)
}
