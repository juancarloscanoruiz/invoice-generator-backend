package controllers

import (
	controllestypes "invoice-app/controllers/types"
	"invoice-app/database"
	"invoice-app/database/models"
	"invoice-app/utils"

	"github.com/gofiber/fiber/v2"
)

func ListInvoices(c *fiber.Ctx) error {
	var invoices []models.Invoice

	if err := database.Db.Preload("Client").Find(&invoices).Error; err != nil {
		errorMessage := err.Error()
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorMessage})
	}

	response := utils.APIResponse{Status: fiber.StatusOK, Data: invoices, Error: nil}
	return c.JSON(response)

}

func CreateInvoice(c *fiber.Ctx) error {
	var body controllestypes.CreateInvoiceRequest
	c.BodyParser(&body)

	invoice := models.Invoice{
		StreetAddress:      body.StreetAddress,
		City:               body.City,
		PostCode:           body.PostCode,
		Country:            body.Country,
		Status:             body.Status,
		IssueDate:          body.IssueDate,
		PaymentTerms:       body.PaymentTerms,
		ProjectDescription: body.ProjectDescription,
		Client:             body.Client,
	}
	var items []models.Item

	database.Db.Create(&invoice)
	for _, item := range body.Items {
		item.InvoiceID = invoice.ID
		items = append(items, item)
	}
	database.Db.Create(&items)
	return c.JSON(items)
}
