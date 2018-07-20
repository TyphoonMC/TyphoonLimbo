package main

import (
	"fmt"
	"github.com/TyphoonMC/TyphoonCore"
)

func main() {
	core := typhoon.Init()
	core.SetBrand("typhoonlimbo")

	loadConfig(core)

	core.On(func(e *typhoon.PlayerJoinEvent) {
		if &joinMessage != nil {
			e.Player.WritePacket(&joinMessage)
		}
		if &bossbarCreate != nil {
			e.Player.WritePacket(&bossbarCreate)
		}
		if &playerListHF != nil {
			e.Player.WritePacket(&playerListHF)
		}
	})

	core.On(func(e *typhoon.PlayerChatEvent) {
		e.Player.WritePacket(&typhoon.PacketPlayMessage{
			Component: fmt.Sprintf(`{"text":"<%s> %s"}`, e.Player.GetName(), typhoon.JsonEscape(e.Message)),
			Position:  typhoon.CHAT_BOX,
		})
	})

	core.Start()
}
