package models

type TravelItem struct {
	Id                 int     `json:"id" xorm:"'id' pk autoincr"`
	Total_price        float64 `xorm:"not null 'total_price'"`
	Itinerary          string  `xorm:"not null 'itinerary'"`
	Traveler_name      string  `xorm:"not null 'traveler_name'"`
	Ticket_number      int     `xorm:"not null 'ticket_number'"`
	Conjunction_number string  `xorm:"not null 'conjunction_number'"`
	Transaction_type   string  `xorm:"not null 'transaction_type'"`
	Product_type       string  `xorm:"not null 'product_type'"`
	Status             string  `xorm:"not null 'status'"`
	Id_invoice         int     `xorm:"'id_invoice'"`
}

func (*TravelItem) TableName() string {
	return "air_booking"
}
