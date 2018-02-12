package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/TyphoonMC/go.uuid"
	"io/ioutil"
)

type Config struct {
	ListenAddress    string          `json:"listen_address"`
	MaxPlayers       int             `json:"max_players"`
	Motd             string          `json:"motd"`
	Restricted       bool            `json:"restricted"`
	Logs             bool            `json:"logs"`
	JoinMessage      json.RawMessage `json:"join_message"`
	BossBar          json.RawMessage `json:"boss_bar"`
	PlayerListHeader json.RawMessage `json:"playerlist_header"`
	PlayerListFooter json.RawMessage `json:"playerlist_footer"`
	Compression      bool            `json:"enable_compression"`
	Threshold        int             `json:"compression_threshold"`
}

var (
	config         Config
	join_message   PacketPlayMessage
	bossbar_create PacketBossBar
	playerlist_hf  PacketPlayerListHeaderFooter
	favicon        string
)

func InitConfig() (err error) {
	fav, err := ioutil.ReadFile("./favicon.png")
	if err == nil {
		favicon = "data:image/png;base64," + base64.StdEncoding.EncodeToString(fav)
	}

	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	if config.JoinMessage != nil {
		join_message = PacketPlayMessage{
			string(config.JoinMessage),
			CHAT_BOX,
		}
	}

	if config.BossBar != nil {
		bossbar_create = PacketBossBar{
			uuid:     uuid.Must(uuid.NewV4()),
			action:   BOSSBAR_ADD,
			title:    string(config.BossBar),
			health:   1.0,
			color:    BOSSBAR_COLOR_RED,
			division: BOSSBAR_NODIVISION,
			flags:    0,
		}
	}

	playerlist_hf = PacketPlayerListHeaderFooter{}
	if config.PlayerListHeader != nil {
		msg := string(config.PlayerListHeader)
		playerlist_hf.header = &msg
	}
	if config.PlayerListFooter != nil {
		msg := string(config.PlayerListFooter)
		playerlist_hf.footer = &msg
	}
	return
}
