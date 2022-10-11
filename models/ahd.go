package models

type AH struct {
	No    int    `form:"no" json:"no"`
	Arab  string `form:"arab" json:"arab"`
	Latin string `form:"latin" json:"latin"`
	Arti  string `form:"arti" json:"arti"`
}

type R_ah struct {
	Data []AH
}
