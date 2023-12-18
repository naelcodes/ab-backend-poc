// models/index.go
package models

type Customer struct {
	Id                int    `json:"id" xorm:"'id' pk autoincr"`
	Customer_name     string `xorm:"not null 'customer_name'"`
	Account_number    string `xorm:"not null 'account_number'"`
	Id_currency       int    `xorm:"'id_currency'"`
	Id_country        int    `xorm:"'id_country'"`
	Alias             string `xorm:"not null unique 'alias'"`
	Ab_key            string `xorm:"not null unique 'ab_key'"`
	State             string `xorm:"not null 'state'"`
	Tmc_client_number string `xorm:"not null unique 'tmc_client_number'"`
	Tag               string `xorm:" not null 'tag' "`
}

func (*Customer) TableName() string {
	return "customer"
}
