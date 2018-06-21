package pbgame

type MsgType = uint8

//main type
const (
	MainSystem  MsgType = 0
	MainAccount MsgType = 1
	MainGame    MsgType = 2
)

//sub type
const (
	SubLoginReq = 1
	SubLoginRsp = 2
)

const (
	SubEnterGameReq = 1
	SubEnterGameRsp = 2
	SubRollReq      = 3
	SubRollRsp      = 4
)
