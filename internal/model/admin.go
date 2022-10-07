package model

type Admin struct {
	Id       uint   `json:"id"`
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
