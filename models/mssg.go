package models

import "time"

type Mail struct {
	Datetime time.Time `gorm:"datetime" json:"datetime"`
	Id       int64     `json:"id"`
	Subject  string    `json:"subjectS"`
	Message  string    `json:"message"`
}
