package models

type Invoice struct {
	Id            int     `xorm:"'id' pk autoincr" json:"id"`
	CreationDate  string  `xorm:"not null 'creation_date'" json:"creationDate"`
	InvoiceNumber string  `xorm:"not null 'invoice_number'" json:"invoiceNumber"`
	DueDate       string  `xorm:"not null 'due_date'" json:"dueDate"`
	Status        string  `xorm:"not null 'status'" json:"status"`
	Amount        float64 `xorm:"not null 'amount'" json:"amount"`
	Balance       float64 `xorm:"not null 'balance'" json:"balance"`
	NetAmount     float64 `xorm:"not null 'net_amount'" json:"-"`
	BaseAmount    float64 `xorm:"not null 'base_amount'" json:"-"`
	CreditApply   float64 `xorm:"not null 'credit_apply'" json:"creditApply"`
	Tag           string  `xorm:" not null 'tag' " json:"-"`

	IdCustomer int `xorm:"'id_customer'" json:"idCustomer,omitempty"`
}

func (*Invoice) TableName() string {
	return "invoice"
}
