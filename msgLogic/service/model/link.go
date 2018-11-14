package model

import "time"

type Link struct {
	Id      string
	Key     string
	AppId   string
	Created time.Time
	Updated time.Time
}
