package models

type Ts struct {
	Id          int    `json:"id"`
	Name        string ` json:"name"`
	Price       string ` json:"price "`
	Image       string ` json:"image"`
	Description string ` json:"description"`
}

type Test struct {
	Data []Ts
}
