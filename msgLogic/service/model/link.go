package model

import "time"

type Link struct {
	Id      string    `json:"id"`
	Key     string    `json:"key"`
	Nick    string    `json:"nick"`
	Avt     string    `json:"avt"`
	AppId   string    `json:"appId"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
