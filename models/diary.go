package models

import "time"

type Diary struct {
	No    int64     `gorm:"primaryKey" json:"no"`
	Time  time.Time `gorm:"datetime" json:"time"`
	User  int64     `json:"user"`
	Judul string    `json:"judul"`
	Isi   string    `json:"isi"`
}

type Diaryscn struct {
	No    int64  `gorm:"primaryKey" json:"no"`
	Time  string `json:"time"`
	Judul string `json:"judul"`
}

type Dd struct {
	Data []Diaryscn
}

type Diarys struct {
	No    int64  `gorm:"primaryKey" json:"no"`
	Time  string `json:"time"`
	Judul string `json:"judul"`
	Isi   string `json:"isi"`
}

type Dd2 struct {
	Data []Diarys
}
