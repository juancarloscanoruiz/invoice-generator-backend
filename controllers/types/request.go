package types

import "invoice-app/database/models"

type CreateInvoiceRequest struct {
	models.Invoice
	Items []models.Item
}
