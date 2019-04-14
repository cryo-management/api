package common

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Active    string `json:"active"`
}

type env struct {
	User user `json:"user"`
}

var Session env
