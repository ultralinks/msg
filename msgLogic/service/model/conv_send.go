package model

import "time"

type ConvSend struct {
	MsgId    string
	ConvId   string
	ToLinkId string
	Status   int
	Created  time.Time
	Updated  time.Time
}
