package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`
}

func ParseConfigFromJSONFile(filename string) (c *Config, err error) {
	f, err := os.Open(filename)
	if err != nil {

	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)

	return
}
