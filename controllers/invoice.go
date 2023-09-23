package controllers

import (
	"invoice-app/database"
	"invoice-app/database/models"
	"invoice-app/utils"

	"github.com/gofiber/fiber/v2"
)

func ListInvoices(c *fiber.Ctx) error {
	var invoices []models.Invoice
	// if err := database.Db.InnerJoins("clients").Find(&invoices).Error; err != nil {
	if err := database.Db.Joins("INNER JOIN clients on clients.id = invoices.client_id").Find(&invoices).Error; err != nil {
		errorMessage := err.Error()
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorMessage})
	}
	response := utils.APIResponse{Status: fiber.StatusOK, Data: invoices, Error: nil}
	return c.JSON(response)
}

type CreateInvoiceRequest struct {
	Invoice models.Invoice
	Items   []models.Item
}

func CreateInvoice(c *fiber.Ctx) error {
	var body CreateInvoiceRequest
	c.BodyParser(&body)

	invoice := models.Invoice{
		StreetAddress:      body.Invoice.StreetAddress,
		City:               body.Invoice.City,
		PostCode:           body.Invoice.PostCode,
		Country:            body.Invoice.Country,
		Status:             body.Invoice.Status,
		IssueDate:          body.Invoice.IssueDate,
		PaymentTerms:       body.Invoice.PaymentTerms,
		ProjectDescription: body.Invoice.ProjectDescription,
		Client:             body.Invoice.Client,
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
