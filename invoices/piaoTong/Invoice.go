package piaoTong

import "invoiceOcr/invoices"

type Invoice struct {
	invoices.InvoiceChannel
}

func (i Invoice) CheckParams(p *invoices.Context) error {
	panic("implement me")
}

func (i Invoice) InvoiceByVoucher(p *invoices.Context) error {
	panic("implement me")
}

func (i Invoice) Run(p *invoices.Context) {

}
