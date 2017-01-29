package main

var (
	players map[int]*Player = make(map[int]*Player)
)

func (player *Player) register(id int) {
	players[id]=player
}

func (player *Player) unregister(id int) {
	delete(players, id)
}