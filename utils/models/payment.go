package models

type Payment struct {
	Id                   int     `xorm:"'id' pk autoincr"`
	Payment_number       string  `xorm:"not null 'number'"`
	Payment_date         string  `xorm:"not null 'date'"`
	Balance              float64 `xorm:"not null 'balance'"`
	Amount               float64 `xorm:"not null 'amount'"`
	Base_amount          float64 `xorm:"not null 'base_amount'"`
	Used_amount          float64 `xorm:"not null 'used_amount'"`
	Type                 string  `xorm:"not null 'type'"`
	Fop                  string  `xorm:"not null 'fop'"`
	Status               string  `xorm:"not null 'status'"`
	Id_chart_of_accounts int     `xorm:"not null 'id_chart_of_accounts'"`
	Id_currency          int     `xorm:"not null 'id_currency'"`
	Id_customer          *int    `xorm:"'id_customer'"`
	Customer             *string `xorm:" <- 'customer' "`
	Tag                  string  `xorm:" not null 'tag' "`
}

func (*Payment) TableName() string {
	return "payment_received"
}
