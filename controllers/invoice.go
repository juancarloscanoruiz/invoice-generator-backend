package controllers

import (
	controllestypes "invoice-app/controllers/types"
	"invoice-app/database"
	"invoice-app/database/models"
	"invoice-app/database/repositories"
	"invoice-app/utils"

	"github.com/gofiber/fiber/v2"
)

func ListInvoices(c *fiber.Ctx) error {

	invoices, err := repositories.GetInvoices()
	if err != nil {
		errorMessage := err.Error()
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorMessage})
	}
	response := utils.APIResponse{Status: fiber.StatusOK, Data: invoices, Error: nil}
	return c.JSON(response)

}

func CreateInvoice(c *fiber.Ctx) error {
	var body controllestypes.CreateInvoiceRequest

	bodyParserErr := c.BodyParser(&body)
	var errorResponseMessage string

	if bodyParserErr != nil {
		errorResponseMessage = "Missing required data for creating a new invoice. Please provide the necessary information in the request body."
		return c.JSON(utils.APIResponse{Status: fiber.StatusBadRequest, Data: nil, Error: &errorResponseMessage})
	}

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
	invoiceErr := database.Db.Create(&invoice).Error

	var items []models.Item
	for _, item := range body.Items {
		item.InvoiceID = invoice.ID
		items = append(items, item)
	}

	itemErr := database.Db.Create(&items).Error
	if itemErr != nil || invoiceErr != nil {
		errorResponseMessage = "There was an error creating the invoice."
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorResponseMessage})
	}

	return c.JSON(utils.APIResponse{Status: fiber.StatusOK, Data: nil, Error: nil})
}
