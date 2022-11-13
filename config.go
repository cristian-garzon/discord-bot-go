package main

import (
	"discord-bot-go/discord"
	"discord-bot-go/environment"
)

type Config struct {
	Discord discord.Config
}

func NewConfig() (*Config, error) {

	discordConfig, err := discord.NewConfig(
		environment.GetEnv("TOKEN", ""),
		environment.GetEnv("PREFIX", ""),
	)

	if err != nil {
		return nil, err
	}

	return &Config{
		Discord: *discordConfig,
	}, nil

}
