package model

import "time"

type Msg struct {
	Id         string
	Key        string
	Data       string
	FromLinkId string
	Created    time.Time
	Updated    time.Time
}
