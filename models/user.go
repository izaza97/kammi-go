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
	Username string `json:"username"`
	Email    string `json:"email"`
	Img      string `json:"img"`
}

type Uf struct {
	Id  int64  `gorm:"primaryKey" json:"id"`
	Img string `json:"img"`
}

//update username
type UU struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
}

//update pw
type UP struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Password string `json:"password"`
}

type UserS struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Img      string `json:"img"`
}

type Un struct {
	Data []UU
}

type Up struct {
	Data []UP
}

type Ud struct {
	Data []UserS
}

type UI struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Status  int64  `json:"status"`
	Message string `json:"Message"`
}

type Imgp struct {
	Id  int    `gorm:"primaryKey" json:"id"`
	Img string `json:"img"`
}

type Ip struct {
	Data []Imgp
}
