package model

import "time"

type Link struct {
	Id      string
	Key     string
	Nick    string
	Avt     string
	AppId   string
	Created time.Time
	Updated time.Time
}
