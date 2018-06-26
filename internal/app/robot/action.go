package robot

import (
	"starter-kit/internal/pkg/pb/pbgame"
)

func Login(robot *Robot) {
	req := &pbgame.LoginRequest{
		Account:  robot.account,
		Password: robot.password,
	}

	robot.SendMsg(pbgame.MainAccount, pbgame.SubLoginReq, req)
}

func Rolll(robot *Robot) {
	req := &pbgame.RollRequest{}
	robot.SendMsg(pbgame.MainGame, pbgame.SubRollReq, req)
}
