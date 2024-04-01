// models/index.go
package models

type Customer struct {
	Id           int    `xorm:"'id' pk autoincr" json:"id,omitempty"`
	CustomerName string `xorm:"not null 'customer_name'" json:"customerName,omitempty"`
	// Street             string `xorm:"not null 'street'" json:"street,omitempty"`
	State string `xorm:"not null 'state'" json:"state,omitempty"`
	// ZipCode            string `xorm:"not null 'zip_code'" json:"zipCode,omitempty"`
	// Notes              string `xorm:"not null 'notes'" json:"notes,omitempty"`
	// Terms              int    `xorm:"'terms'" json:"terms,omitempty"`
	AccountNumber string `xorm:"not null 'account_number'" json:"accountNumber,omitempty"`
	// TaxId              string `xorm:"not null 'tax_id'" json:"taxId,omitempty"`
	// Balance            string `xorm:"'balance'" json:"balance,omitempty"`
	// CreditLimit        string `xorm:"'credit_limit'" json:"creditLimit,omitempty"`
	// IsActive           bool   `xorm:"'is_active'" json:"isActive,omitempty"`
	// IsSubAgency        bool   `xorm:"'is_sub_agency'" json:"isSubAgency,omitempty"`
	// OpeningBalance     string `xorm:"'opening_balance'" json:"openingBalance,omitempty"`
	// Language           string `xorm:"'language'" json:"language,omitempty"`
	IdCurrency int `xorm:"'id_currency'" json:"idCurrency,omitempty"`
	IdCountry  int `xorm:"'id_country'" json:"idCountry,omitempty"`
	// IrsShareKey        string `xorm:"'irs_share_key'" json:"irsShareKey,omitempty"`
	// Agency             string `xorm:"'agency'" json:"agency,omitempty"`
	// OpeningBalanceDate string `xorm:"'opening_balance_date'" json:"openingBalanceDate,omitempty"`
	// AvoidDeletion      bool   `xorm:"'avoid_deletion'" json:"avoidDeletion,omitempty"`
	// IsEditable         bool   `xorm:"'is_editable'" json:"isEditable,omitempty"`
	Alias string `xorm:"unique 'alias'" json:"alias,omitempty"`
	// AlreadyUsed        bool   `xorm:"'already_used'" json:"alreadyUsed,omitempty"`
	AbKey           string `xorm:"unique 'ab_key'" json:"abKey,omitempty"`
	TmcClientNumber string `xorm:"unique 'tmc_client_number'" json:"tmcClientNumber,omitempty"`
	// Category           string `xorm:"'category'" json:"category,omitempty"` // ab_customer_category
	// TradeRegister      string `xorm:"'trade_register'" json:"tradeRegister,omitempty"`
	// GeneralAccount     string `xorm:"'general_account'" json:"generalAccount,omitempty"`
	// MiscField1         string `xorm:"'misc_field1'" json:"miscField1,omitempty"`
	// MiscField2         string `xorm:"'misc_field2'" json:"miscField2,omitempty"`
	// MiscField3         string `xorm:"'misc_field3'" json:"miscField3,omitempty"`
	// CustomerCode       string `xorm:"'customer_code'" json:"customerCode,omitempty"`
	// IdGroup int `xorm:"'id_group'" json:"idGroup,omitempty"`
}

func (*Customer) TableName() string {
	return "customer"
}
