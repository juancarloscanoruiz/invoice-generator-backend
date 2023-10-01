package controllers

import (
	controllerstypes "invoice-app/controllers/types"
	"invoice-app/database/models"
	"invoice-app/database/repositories"
	"invoice-app/utils"

	"github.com/gofiber/fiber/v2"
)

func ListInvoices(c *fiber.Ctx) error {

	invoices, err := repositories.GetInvoices()
	if err != nil {
		errorMessage := err.Error()
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorMessage})
	}

	response := utils.APIResponse{Status: fiber.StatusOK, Data: invoices, Error: nil}
	c.Status(fiber.StatusOK)
	return c.JSON(response)

}

func CreateInvoice(c *fiber.Ctx) error {
	var body controllerstypes.CreateInvoiceRequest

	bodyParserErr := c.BodyParser(&body)
	var errorResponseMessage string

	if bodyParserErr != nil {
		c.Status(fiber.StatusBadRequest)
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
		Items:              body.Items,
	}

	createdInvoice, createInvoiceErr := repositories.CreateInvoice(invoice)

	if createInvoiceErr != nil {
		errorResponseMessage = "There was an error creating the invoice."
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(utils.APIResponse{Status: fiber.StatusInternalServerError, Data: nil, Error: &errorResponseMessage})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(utils.APIResponse{Status: fiber.StatusOK, Data: createdInvoice, Error: nil})
}

func DeleteInvoice(c *fiber.Ctx) {
	/*
		 	invoiceId, err := strconv.Atoi(c.Params("id"))
			if err != nil {
				c.Status(fiber.StatusBadRequest)

			}
			repositories.DeleteInvoice()
	*/
}
