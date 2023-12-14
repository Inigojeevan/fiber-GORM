package models

type User struct {
	Id int `json:"id" gorm:"unique"`
	//CreatedAt time.Time
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
