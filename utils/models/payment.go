package models

type Payment struct {
	Id                int     `xorm:"'id' pk autoincr"  json:"id"`
	PaymentNumber     string  `xorm:"not null 'number'" json:"paymentNumber"`
	PaymentDate       string  `xorm:"not null 'date'" json:"paymentDate"`
	Balance           float64 `xorm:"not null 'balance'" json:"balance"`
	Amount            float64 `xorm:"not null 'amount'" json:"amount"`
	BaseAmount        float64 `xorm:"not null 'base_amount'" json:"-"`
	UsedAmount        float64 `xorm:"not null 'used_amount'" json:"usedAmount"`
	Type              string  `xorm:"not null 'type'" json:"-"`
	PaymentMode       string  `xorm:"not null 'fop'" json:"paymentMode"`
	Status            string  `xorm:"not null 'status'" json:"status"`
	IdChartOfAccounts int     `xorm:"not null 'id_chart_of_accounts'" json:"-"`
	IdCurrency        int     `xorm:"not null 'id_currency'" json:"-"`
	IdCustomer        int     `xorm:"'id_customer'" json:"idCustomer,omitempty"`
	Tag               string  `xorm:"not null 'tag'" json:"-"`
}

func (*Payment) TableName() string {
	return "payment_received"
}
