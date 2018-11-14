package model

import "time"

type Conv struct {
	Id      string
	Key     string
	Name    string
	Created time.Time
	Updated time.Time
}
