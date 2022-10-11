package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Img      string `json:"img"`
}

type UserU struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Img      string `json:"img"`
}

type UserS struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Img      string `json:"img"`
}
