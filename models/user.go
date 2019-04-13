package models

//User docs
type User struct {
	ID        string `json:"id" sql:"id" pk:"true"`
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`
	Email     string `json:"email" sql:"email"`
	Password  string `json:"password" sql:"password"`
	Language  string `json:"language" sql:"language"`
	Active    string `json:"active" sql:"active"`
}
