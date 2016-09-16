package app

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// ConfToml ...
type ConfToml struct {
	SlackInfo     SlackInfo     `toml:"slack-info"`
	GoogleAccount GoogleAccount `toml:"google-account"`
}

// SlackInfo ...
type SlackInfo struct {
	Token     string `toml:"token"`
	ChannelID string `toml:"channel_id"`
}

// GoogleAccount ...
type GoogleAccount struct {
	Name string `toml:"name"`
	Pass string `toml:"pass"`
}

// LoadConfToml from file
func LoadConfToml(path string) (ConfToml, error) {
	var conf ConfToml
	_, err := toml.DecodeFile(path, &conf)
	if err != nil {
		fmt.Printf("%s\n", err)
		return conf, err
	}

	return conf, err
}

func checkConf(conf ConfToml) {
	if conf.SlackInfo.ChannelID == "" {
		fmt.Println("Please enter the channel_id in conf.toml")
	}
}
