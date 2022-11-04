package models

type Brt struct {
	Id       int    `json:"id"`
	Img      string `json:"img"`
	Judul    string `json:"judul"`
	Katagori string `json:"katagori"`
	Isi      string `json:"isi"`
}

type Dbrt struct {
	Data []Brt
}

type Vbrt struct {
	Id       int    `json:"id"`
	Img      string `json:"img"`
	Judul    string `json:"judul"`
	Katagori string `json:"katagori"`
	Isi      string `json:"isi"`
}

type Dvbrt struct {
	Data []Vbrt
}

type Kbrt struct {
	Id       int    `gorm:"primaryKey" json:"id"`
	Katagori string `json:"katagori"`
}

type Kdb struct {
	Data []Kbrt
}
