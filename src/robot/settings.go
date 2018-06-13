package robot

import (
	"encoding/binary"
	"time"
)

var BinaryCoder = binary.LittleEndian

const (
	DefaultSendChanSize = 1000
	PingPeriod = 1 * time.Second
	PongWait = 2 * time.Second
	WriteWait = 1* time.Second
)
