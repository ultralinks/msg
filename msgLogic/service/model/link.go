package model

import "time"

type Link struct {
	Id      int
	AppId   int
	Key     string
	Created time.Time
	Updated time.Time
}
