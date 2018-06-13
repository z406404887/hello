package game

type Player struct {
	account string
	name string
	id uint32
	money int32
}

func NewPlayer(acc string, name string, id uint32, money int32) *Player  {
	return &Player{
		account:acc,
		name:name,
		id:id,
		money:money,
	}
}