package models

//User defines the struct of this object
type User struct {
	ID        string `json:"id" sql:"id" pk:"true"`
	FirstName string `json:"first_name" sql:"first_name"`
	LastName  string `json:"last_name" sql:"last_name"`
	Email     string `json:"email" sql:"email"`
	Password  string `json:"password" sql:"password"`
	Language  string `json:"language" sql:"language"`
	Active    bool   `json:"active" sql:"active"`
}

//GetID returns object primary key
func (u *User) GetID() string {
	return u.ID
}
