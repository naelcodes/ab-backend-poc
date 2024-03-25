package models

type Traveler struct {
	Id                   int    `xorm:"'id' pk autoincr"  json:"id,omitempty"`
	Title                string `xorm:"'title'" json:"title,omitempty"`
	FirstName            string `xorm:"'first_name'" json:"firstName,omitempty"`
	LastName             string `xorm:"'last_name'" json:"lastName,omitempty"`
	DisplayName          string `xorm:"'display_name'" json:"displayName,omitempty"`
	Email                string `xorm:"'email'" json:"email,omitempty"`
	WorkPhone            string `xorm:"'work_phone'" json:"workPhone,omitempty"`
	MobilePhone          string `xorm:"'mobile_phone'" json:"mobilePhone,omitempty"`
	PassPortNumber       string `xorm:"'passport_number'" json:"passportNumber,omitempty"`
	PassPortIssuedOn     string `xorm:"'passport_issued_on'" json:"passportIssuedOn,omitempty"`
	PassPortExpireOn     string `xorm:"'passport_expire_on'" json:"passportExpiredOn,omitempty"`
	IdCitizenship        int    `xorm:"'id_citizenship'" json:"idCitizenship,omitempty"`
	BirthDate            string `xorm:"'birth_date'" json:"birthDate,omitempty"`
	MaritalStatus        string `xorm:"'marital_status'" json:"maritalStatus,omitempty"` //ab_marital_status
	Gender               string `xorm:"'gender'" json:"gender,omitempty"`
	Street               string `xorm:"'street'" json:"street,omitempty"`
	City                 string `xorm:"'city'" json:"city,omitempty"`
	State                string `xorm:"'state'" json:"state,omitempty"`
	PostalCode           string `xorm:"'postal_code'" json:"postalCode,omitempty"`
	IdCountry            int    `xorm:"'id_country'" json:"idCountry,omitempty"`
	Note                 string `xorm:"'note'" json:"note,omitempty"`
	FlightSmoking        bool   `xorm:"'flight_smoking'" json:"flightSmoking,omitempty"`
	FlightSeating        string `xorm:"'flight_seating'" json:"flightSeating,omitempty"`          //ab_seating
	FlightMealRequest    string `xorm:"'flight_meal_request'" json:"flightMealRequest,omitempty"` //ab_meal_request
	FlightSpecialRequest string `xorm:"'flight_special_request'" json:"flightSpecialRequest,omitempty"`
	HotelBedding         string `xorm:"'hotel_bedding'" json:"hotelBedding,omitempty"` //ab_room_type
	HotelSpecialRequest  string `xorm:"'hotel_special_request'" json:"hotelSpecialRequest,omitempty"`
	CarSize              string `xorm:"'car_size'" json:"carSize,omitempty"` //ab_car_size
	CarSpecialRequest    string `xorm:"'car_special_request'" json:"carSpecialRequest,omitempty"`
	IsActive             bool   `xorm:"'is_active'" json:"isActive,omitempty"`
}

func (*Traveler) TableName() string {
	return "traveler"
}
