package main

var (
	players map[int]*Player = make(map[int]*Player)
)

func (player *Player) register() {
	players[player.id]=player
}

func (player *Player) unregister() {
	delete(players, player.id)
}