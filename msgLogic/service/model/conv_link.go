package model

import "time"

type ConvLink struct {
	ConvId   string
	LinkId   string
	IsOwner  int
	IsMute   int
	IsIgnore int
	Created  time.Time
	Updated  time.Time
}
