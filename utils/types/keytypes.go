package types

type KeyType int

const (
	InvoiceId KeyType = iota
	TravelItemIds
	Transaction
)
