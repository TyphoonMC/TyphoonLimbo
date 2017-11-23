package main

import "sync"

var (
	players_count = 0
	players map[int]*Player = make(map[int]*Player)
	playersMutex = &sync.Mutex{}
)

func (player *Player) register() {
	playersMutex.Lock()
	players[player.id]=player
	players_count++
	playersMutex.Unlock()
}

func (player *Player) unregister() {
	playersMutex.Lock()
	if _, ok := players[player.id]; ok {
		players_count--
	}
	delete(players, player.id)
	playersMutex.Unlock()
}
