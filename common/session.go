package common

//User docs
type user struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Language  string `json:"language"`
	Active    string `json:"active"`
}

//env docs
type env struct {
	User user `json:"user"`
}

//Session docs
var Session env
