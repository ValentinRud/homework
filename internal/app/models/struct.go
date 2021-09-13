package models

type AddUser struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Status    string `json:"status"`
}

type SeeUser struct {
	Id       int    `json:"id"`
	LastName string `json:"lastName"`
}
