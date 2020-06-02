package main

import (
	"fmt"
	t "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := t.Init()
	core.SetBrand("typhoonlimbo")
	core.SetGamemode(t.SURVIVAL)

	loadConfig(core)

	if spawn != nil {
		fmt.Println("Using schematic world")
		if config.Spawn.Location != nil {
			spawn.Spawn = *config.Spawn.Location
		}
		core.SetMap(spawn)
	}

	core.On(func(e *t.PlayerJoinEvent) {
		if config.JoinMessage != nil {
			e.Player.SendRawMessage(string(config.JoinMessage))
		}
		if &bossbarCreate != nil {
			e.Player.WritePacket(&bossbarCreate)
		}
		if &playerListHF != nil {
			e.Player.WritePacket(&playerListHF)
		}
	})

	core.On(func(e *t.PlayerChatEvent) {
		msg := t.ChatMessage("")
		msg.SetExtra([]t.IChatComponent{
			t.ChatMessage("<"),
			t.ChatMessage(e.Player.GetName()),
			t.ChatMessage("> "),
			t.ChatMessage(e.Message),
		})
		e.Player.SendMessage(msg)
	})

	core.Start()
}

type ChunkSave struct {
	X       int    `json:"x"`
	Y       int    `json:"y"`
	Bitmask int    `json:"bitmask"`
	Data    string `json:"data"`
}
