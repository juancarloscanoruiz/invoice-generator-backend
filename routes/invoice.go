package routes

import (
	"invoice-app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupInvoceRoutes(app *fiber.App) {
	app.Get("/invoice/:id", controllers.ListInvoices)
	app.Post("/invoice", controllers.CreateInvoice)
	app.Delete("/invoice/:id", controllers.CreateInvoice)
}
