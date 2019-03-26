package main

import (
	"encoding/json"
	t "github.com/TyphoonMC/TyphoonCore"
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
	bossbarCreate t.PacketBossBar
	playerListHF  t.PacketPlayerListHeaderFooter
)

func loadConfig(core *t.Core) {
	core.GetConfig(&config)

	if config.BossBar != nil {
		bossbarCreate = t.PacketBossBar{
			UUID:     uuid.NewV4(),
			Action:   t.BOSSBAR_ADD,
			Title:    string(config.BossBar),
			Health:   1.0,
			Color:    t.BOSSBAR_COLOR_RED,
			Division: t.BOSSBAR_NODIVISION,
			Flags:    0,
		}
	}

	playerListHF = t.PacketPlayerListHeaderFooter{}
	if config.PlayerListHeader != nil {
		msg := string(config.PlayerListHeader)
		playerListHF.Header = &msg
	}
	if config.PlayerListFooter != nil {
		msg := string(config.PlayerListFooter)
		playerListHF.Footer = &msg
	}
}
