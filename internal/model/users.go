package model

type Users struct {
	Id       uint   `json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}
