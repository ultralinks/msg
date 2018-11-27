package model

import "time"

type Conv struct {
	Id      string    `json:"id"`
	Key     string    `json:"key"`
	Name    string    `json:"name"`
	Avt     string    `json:"avt"`
	Type    string    `json:"type"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
