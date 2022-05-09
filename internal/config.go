package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token            string `json:"token"`
	Prefix           string `json:"prefix"`
	Owner            string `json:"owner"`
	OwnerStackTraces string `json:"sendOwnerStackTraces"`
}

func InitConfigFromJSONFile(fileName string) (c *Config, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}

	c = new(Config)
	err = json.NewDecoder(f).Decode(c)

	return
}
