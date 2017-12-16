package main

import (
	"encoding/json"
	"io/ioutil"
	"encoding/base64"
	"github.com/satori/go.uuid"
)

var (
	config map[string]interface{}
	join_message PacketPlayMessage
	bossbar_create PacketBossBar
	playerlist_hf PacketPlayerListHeaderFooter
	compressionEnabled = false
	compressionThreshold = 256
)

func InitConfig() (err error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}

	if config["join_message"] != nil {
		message, err := base64.StdEncoding.DecodeString(config["join_message"].(string))
		if err != nil {
			panic(err)
		}
		join_message = PacketPlayMessage{
			string(message),
			CHAT_BOX,
		}
	}
	if config["boss_bar"] != nil {
		message, err := base64.StdEncoding.DecodeString(config["boss_bar"].(string))
		if err != nil {
			panic(err)
		}
		bossbar_create = PacketBossBar{
			uuid: uuid.NewV4(),
			action: BOSSBAR_ADD,
			title: string(message),
			health: 1.0,
			color: BOSSBAR_COLOR_RED,
			division: BOSSBAR_NODIVISION,
			flags: 0,
		}
	}

	playerlist_hf = PacketPlayerListHeaderFooter{}
	if config["playerlist_header"] != nil {
		dat, err := base64.StdEncoding.DecodeString(config["playerlist_header"].(string))
		if err != nil {
			panic(err)
		}
		msg := string(dat)
		playerlist_hf.header = &msg
	}
	if config["playerlist_footer"] != nil {
		dat, err := base64.StdEncoding.DecodeString(config["playerlist_footer"].(string))
		if err != nil {
			panic(err)
		}
		msg := string(dat)
		playerlist_hf.footer = &msg
	}

	if config["enable_compression"] != nil {
		compressionEnabled = config["enable_compression"].(bool)
	}
	if config["compression_threshold"] != nil {
		compressionThreshold = int(config["compression_threshold"].(float64))
	}
	return
}
