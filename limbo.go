package main

import (
	t "github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := t.Init()
	core.SetBrand("typhoonlimbo")

	loadConfig(core)

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
