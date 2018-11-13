package model

import "time"

type ConvSend struct {
	MsgId    int
	ConvId   int
	ToLinkId int
	Status   int
	Created  time.Time
	Updated  time.Time
}
