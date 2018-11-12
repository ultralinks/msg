package model

import "time"

type Msg struct {
	Id         int
	Data       string
	FromLinkId int
	Created    time.Time
	Updated    time.Time
}
