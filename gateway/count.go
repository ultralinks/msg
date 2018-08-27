package main

import "sync/atomic"

type Count struct {
	Client int64
}

var CountVal = &Count{}

func Client_INCR() {
	atomic.AddInt64(&CountVal.Client, 1)
}

func Client_DESC() {
	atomic.AddInt64(&CountVal.Client, -1)
}
