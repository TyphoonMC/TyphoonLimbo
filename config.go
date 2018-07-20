package main

import (
	"encoding/json"
	"github.com/TyphoonMC/TyphoonCore"
	"github.com/TyphoonMC/go.uuid"
)

type Config struct {
	JoinMessage      json.RawMessage `json:"join_message"`
	BossBar          json.RawMessage `json:"boss_bar"`
	PlayerListHeader json.RawMessage `json:"playerlist_header"`
	PlayerListFooter json.RawMessage `json:"playerlist_footer"`
}

var (
	config        Config
	joinMessage   typhoon.PacketPlayMessage
	bossbarCreate typhoon.PacketBossBar
	playerListHF  typhoon.PacketPlayerListHeaderFooter
)

func loadConfig(core *typhoon.Core) {
	core.GetConfig(&config)

	if config.JoinMessage != nil {
		joinMessage = typhoon.PacketPlayMessage{
			string(config.JoinMessage),
			typhoon.CHAT_BOX,
		}
	}

	if config.BossBar != nil {
		bossbarCreate = typhoon.PacketBossBar{
			uuid.Must(uuid.NewV4()),
			typhoon.BOSSBAR_ADD,
			string(config.BossBar),
			1.0,
			typhoon.BOSSBAR_COLOR_RED,
			typhoon.BOSSBAR_NODIVISION,
			0,
		}
	}

	playerListHF = typhoon.PacketPlayerListHeaderFooter{}
	if config.PlayerListHeader != nil {
		msg := string(config.PlayerListHeader)
		playerListHF.Header = &msg
	}
	if config.PlayerListFooter != nil {
		msg := string(config.PlayerListFooter)
		playerListHF.Footer = &msg
	}
}
