package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	config map[string]interface{}
)

func InitConfig() (err error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	return
}