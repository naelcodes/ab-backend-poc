// models/index.go
package models

type Customer struct {
	Id              int    `xorm:"'id' pk autoincr" json:"id"`
	CustomerName    string `xorm:"not null 'customer_name'" json:"customerName"`
	AccountNumber   string `xorm:"not null 'account_number'" json:"accountNumber"`
	IdCurrency      int    `xorm:"'id_currency'" json:"-"`
	IdCountry       int    `xorm:"'id_country'" json:"-"`
	Alias           string `xorm:"not null unique 'alias'" json:"alias"`
	AbKey           string `xorm:"not null unique 'ab_key'" json:"abKey"`
	State           string `xorm:"not null 'state'" json:"state"`
	TmcClientNumber string `xorm:"not null unique 'tmc_client_number'" json:"tmcClientNumber"`
	Tag             string `xorm:" not null 'tag' " json:"-"`
}

func (*Customer) TableName() string {
	return "customer"
}
