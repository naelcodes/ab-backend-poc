// models/index.go
package models

type Customer struct {
	Id                int    `json:"id" xorm:"'id' pk autoincr"`
	Customer_name     string `json:"customer_name" xorm:"not null 'customer_name'"`
	Account_number    string `json:"account_number" xorm:"not null 'account_number'"`
	Id_currency       int    `json:"id_currency" xorm:"'id_currency'"`
	Id_country        int    `json:"id_country" xorm:"'id_country'"`
	Alias             string `json:"alias" xorm:"not null unique 'alias'"`
	Ab_key            string `json:"ab_key" xorm:"not null unique 'ab_key'"`
	State             string `json:"state" xorm:"not null 'state'"`
	Tmc_client_number string `json:"tmc_client_number" xorm:"not null unique 'tmc_client_number'"`
	Tag               string `json:"tag" xorm:" not null 'tag' "`
}

func (*Customer) TableName() string {
	return "customer"
}
