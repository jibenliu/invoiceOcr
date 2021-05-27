package baiWang

import (
	invoice "invoiceOcr/services/invoices"
)

type InputData struct {
	invoice.InvoiceChannel
}

func (i InputData) CheckParams(p *invoice.Context) error {
	return invoice.ErrorConfigEmpty
}

func (i InputData) InvoiceByVoucher(p *invoice.Context) error {
	return invoice.ErrorMethodNotAllowed
}

func (i InputData) Run(p *invoice.Context) error {
	return invoice.ErrorServiceNotAvailable
}
