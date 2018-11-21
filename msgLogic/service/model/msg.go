package model

import "time"

type Msg struct {
	Id         string    `json:"id"`
	Key        string    `json:"key"`
	Data       string    `json:"data"`
	FromLinkId string    `json:"fromLinkId"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}
