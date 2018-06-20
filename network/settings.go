package network

import (
	"encoding/binary"
	"time"
)

var BinaryCoder = binary.LittleEndian

const (
	DefaultSendChanSize = 10000
	PingPeriod          = 54 * time.Second
	PongWait            = 60 * time.Second
	WriteWait           = 10 * time.Second
)
