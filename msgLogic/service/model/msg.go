package model

import "time"

type Msg struct {
	Id          string    `json:"id"`
	Key         string    `json:"key"`
	Data        string    `json:"data"`
	FromLinkKey string    `json:"fromLinkKey"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}
