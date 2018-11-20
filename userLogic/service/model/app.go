package model

import "time"

type App struct {
	Id      string
	OrgId   string
	Key     string
	Secret  string
	Name    string
	Desc    string
	Created time.Time
	Updated time.Time
}
