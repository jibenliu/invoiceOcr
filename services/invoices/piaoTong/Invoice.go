package piaoTong

import (
	invoices2 "invoiceOcr/services/invoices"
)

type Invoice struct {
	invoices2.InvoiceChannel
}

func (i Invoice) CheckParams(p *invoices2.Context) error {
	panic("implement me")
}

func (i Invoice) InvoiceByVoucher(p *invoices2.Context) error {
	panic("implement me")
}

func (i Invoice) Run(p *invoices2.Context) {

}
