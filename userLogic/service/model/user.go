package model

import "time"

type User struct {
	Id       string
	Nick     string
	Email    string
	Phone    string
	Password string
	Avt      string
	Status   string
	Logined  time.Time
	Created  time.Time
	Updated  time.Time
}
