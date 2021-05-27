package invoices

type InvoiceType int

const (
	VatSpecialInvoice           InvoiceType = 01
	VatNormalInvoice                        = 04
	VatElectronicNormalInvoice              = 10
	VatElectronicSpecialInvoice             = 20 //待定
)

type Invoice struct {
	serialNo             string        `gorm:"column:serial_no"`
	InvoiceType          int           `gorm:"column:invoice_type" json:"invoice_type" form:"invoice_type"`
	InvoiceNo            string        `gorm:"column:invoice_no" json:"invoice_no" form:"invoice_no"`
	InvoiceCode          string        `gorm:"column:invoice_code" json:"invoice_code" form:"invoice_code"`
	SellerTaxCode        string        `gorm:"column:seller_tax_code" json:"seller_tax_code" form:"seller_tax_code"`
	SellerName           string        `gorm:"column:seller_name" json:"seller_name" form:"seller_name"`
	SellerAddress        string        `gorm:"column:seller_address" json:"seller_address" form:"seller_address"`
	SellerTel            string        `gorm:"column:seller_tel" json:"seller_tel" form:"seller_tel"`
	SellerBankName       string        `gorm:"column:seller_bank_name" json:"seller_bank_name" form:"seller_bank_name"`
	SellerBankAccount    string        `gorm:"column:seller_bank_account" json:"seller_bank_account" form:"seller_bank_account"`
	BuyerTaxCode         string        `gorm:"column:buyer_tax_code" json:"buyer_tax_code" form:"buyer_tax_code"`
	BuyerName            string        `gorm:"column:buyer_name" json:"buyer_name" form:"buyer_name"`
	BuyerAddress         string        `gorm:"column:buyer_address" json:"buyer_address" form:"buyer_address"`
	BuyerTel             string        `gorm:"column:buyer_tel" json:"buyer_tel" form:"buyer_tel"`
	BuyerBankName        string        `gorm:"column:buyer_bank_name" json:"buyer_bank_name" form:"buyer_bank_name"`
	BuyerBankAccount     string        `gorm:"column:buyer_bank_account" json:"buyer_bank_account" form:"buyer_bank_account"`
	AmountWithTax        float64       `gorm:"column:amount_with_tax" json:"amount_with_tax" form:"amount_with_tax"`
	AmountWithoutTax     float64       `gorm:"column:amount_without_tax" json:"amount_without_tax" form:"amount_without_tax"`
	CapitalAmountWithTax string        `gorm:"column:capital_amount_with_tax" json:"capital_amount_with_tax" form:"capital_amount_with_tax"`
	TaxAmount            float64       `gorm:"column:tax_amount" json:"tax_amount" form:"tax_amount"`
	Deductions           float64       `gorm:"column:deductions" json:"deductions" form:"deductions"`
	Remark               string        `gorm:"column:remark" json:"remark" form:"remark"`
	IsValid              int64         `gorm:"column:is_valid" json:"is_valid" form:"is_valid"`
	Payee                string        `gorm:"column:payee" json:"payee" form:"payee"`
	Checker              string        `gorm:"column:checker" json:"checker" form:"checker"`
	Drawer               string        `gorm:"column:drawer" json:"drawer" form:"drawer"`
	InvoiceDate          int64         `gorm:"column:invoice_date" json:"invoice_date" form:"invoice_date"`
	CheckCode            string        `gorm:"column:check_code" json:"check_code" form:"check_code"`
	MachineCode          string        `gorm:"column:machine_code" json:"machine_code" form:"machine_code"`
	VerifyTimes          int64         `gorm:"column:verify_times" json:"verify_times" form:"verify_times"`
	Flag                 int64         `gorm:"column:flag" json:"flag" form:"flag"`
	Receiver             string        `gorm:"column:receiver" json:"receiver" form:"receiver"`
	ReceiverPhone        string        `gorm:"column:receiver_phone" json:"receiver_phone" form:"receiver_phone"`
	ReceiverEmail        string        `gorm:"column:receiver_email" json:"receiver_email" form:"receiver_email"`
	Details              InvoiceDetail `json:"details"`
}

type InvoiceDetail struct {
	ItemName         string  `gorm:"column:item_name" json:"item_name" form:"item_name"`
	Specification    string  `gorm:"column:specification" json:"specification" form:"specification"`
	QuantityUnit     string  `gorm:"column:quantity_unit" json:"quantity_unit" form:"quantity_unit"`
	Quantity         float64 `gorm:"column:quantity" json:"quantity" form:"quantity"`
	PriceType        int64   `gorm:"column:price_type" json:"price_type" form:"price_type"`
	TaxRate          float64 `gorm:"column:tax_rate" json:"tax_rate" form:"tax_rate"`
	Price            float64 `gorm:"column:price" json:"price" form:"price"`
	AmountWithTax    float64 `gorm:"column:amount_with_tax" json:"amount_with_tax" form:"amount_with_tax"`
	AmountWithoutTax float64 `gorm:"column:amount_without_tax" json:"amount_without_tax" form:"amount_without_tax"`
	TaxAmount        float64 `gorm:"column:tax_amount" json:"tax_amount" form:"tax_amount"`
	Deductions       float64 `gorm:"column:deductions" json:"deductions" form:"deductions"`
	GoodsVersion     string  `gorm:"column:goods_version" json:"goods_version" form:"goods_version"`
	GoodsTaxNo       string  `gorm:"column:goods_tax_no" json:"goods_tax_no" form:"goods_tax_no"`
	GoodsTaxName     string  `gorm:"column:goods_tax_name" json:"goods_tax_name" form:"goods_tax_name"`
	TaxPreCon        string  `gorm:"column:tax_pre_con" json:"tax_pre_con" form:"tax_pre_con"`
}

type Config struct {
}
