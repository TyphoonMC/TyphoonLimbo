package main

var (
	players_count = 0
	players map[int]*Player = make(map[int]*Player)
)

func (player *Player) register() {
	players[player.id]=player
	players_count++
}

func (player *Player) unregister() {
	if _, ok := players[player.id]; ok {
		players_count--
	}
	delete(players, player.id)
}