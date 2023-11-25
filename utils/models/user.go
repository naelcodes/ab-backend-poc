// models/index.go
package models

type User struct {
	ID    int    `json:"id" xorm:"'id' pk autoincr"`
	Name  string `json:"name" xorm:"'name'"`
	Email string `json:"email" xorm:"'email'"`
}

func (*User) TableName() string {
	return "user"
}
