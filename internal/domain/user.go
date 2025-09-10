package domain


type User struct {
	ID int `json:"id" sql:"id"`
	FirstName string `json:"firstName" sql:"firstName"`
	LastName string `json:"lastName" sql:"lastName"`
}
