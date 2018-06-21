package game

type Room struct {
	playerMap map[uint32]*Player
}

func NewRoom() *Room {
	return &Room{
		playerMap: make(map[uint32]*Player),
	}
}

func (room *Room) AddPlayer(player *Player) {
	room.playerMap[player.id] = player
}

func (room *Room) RemovePlayer(id uint32) {
	delete(room.playerMap, id)
}
