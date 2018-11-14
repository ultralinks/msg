package model

import "time"

type Conv struct {
	Id      string
	Name    string
	Key     string
	Created time.Time
	Updated time.Time
}
