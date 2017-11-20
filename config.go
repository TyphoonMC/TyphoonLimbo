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
		message, _ := base64.StdEncoding.DecodeString(config["join_message"].(string))
		join_message = PacketPlayMessage{
			string(message),
			CHAT_BOX,
		}
	}
	if config["boss_bar"] != nil {
		message, _ := base64.StdEncoding.DecodeString(config["boss_bar"].(string))
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
	return
}
