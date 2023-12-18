package models

type Invoice struct {
	Id             int     `xorm:"'id' pk autoincr"`
	Creation_date  string  `xorm:"not null 'creation_date'"`
	Invoice_number string  `xorm:"not null 'invoice_number'"`
	Due_date       string  `xorm:"not null 'due_date'"`
	Status         string  `xorm:"not null 'status'"`
	Amount         float64 `xorm:"not null 'amount'"`
	Balance        float64 `xorm:"not null 'balance'"`
	Net_amount     float64 `xorm:"not null 'net_amount'"`
	Base_amount    float64 `xorm:"not null 'base_amount'"`
	Credit_apply   float64 `xorm:"not null 'credit_apply'"`
	Tag            string  `xorm:" not null 'tag' "`

	Id_customer  *int         `xorm:"'id_customer'"`
	Customer     *string      `xorm:"'customer' "`
	Travel_items []TravelItem `xorm:"'travel_items'"`
}

func (*Invoice) TableName() string {
	return "invoice"
}
