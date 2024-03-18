package models

type FacebookUser struct {
	Id        string `json:"id" facebook:"id,required"`
	Email     string `json:"email" facebook:"email"`
	FirstName string `json:"firstname" facebook:"first_name"`
	LastName  string `json:"lastname" facebook:"last_name"`
}

type GoogleUser struct {
	Id            string `json:"sub"`
	Email         string `json:"email"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	EmailVerified bool   `json:"email_verified"`
}
