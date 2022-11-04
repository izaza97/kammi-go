package models

type Dh struct {
	No       int    `json:"no"`
	Katagori int    `json:"katagori"`
	Judul    string `json:"judul"`
	Arab     string `json:"arab"`
	Latin    string `json:"latin"`
	Arti     string `json:"arti"`
}

type Ddh struct {
	Data []Dh
}

type Kdh struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Katagori string `json:"katagori"`
}

type Ddh2 struct {
	Data []Kdh
}
