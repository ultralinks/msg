package model

import "time"

type Org struct {
	Id      string
	Name    string
	Desc    string
	Created time.Time
	Updated time.Time
}
