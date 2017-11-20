package main

import (
	"encoding/json"
	"io/ioutil"
	"encoding/base64"
)

var (
	config map[string]interface{}
	join_message PacketPlayMessage
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
	return
}
