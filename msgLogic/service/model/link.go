package model

import "time"

type Link struct {
	Id      string
	AppId   string
	Key     string
	Created time.Time
	Updated time.Time
}
