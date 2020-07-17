package config

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStruct struct {
	Secret			string	`json:"secret"`
	EffectivTime	int		`json:"effectiv_time"`
}

var Config ConfigStruct

func init() {
	configFile, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configFile, &Config)
	if err != nil {
		panic(err)
	}
}
