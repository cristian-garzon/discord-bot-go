package config

import (
	"os"
	"strings"
)

type AppConfig struct {
	*DiscordConfig
}

func LoadConfig() (*AppConfig, error) {
	discordConfig := &DiscordConfig{
		AuthToken:   getEnvOrDefault("DISCORD_AUTH_TOKEN", "default-token"),
		Description: getEnvOrDefault("DISCORD_DESCRIPTION", "music-bot"),
	}

	return &AppConfig{
		discordConfig,
	}, nil
}

type DiscordConfig struct {
	AuthToken   string
	Description string
}

func getEnvOrDefault(envName string, defaultValue string) string {
	value := strings.TrimSpace(os.Getenv(envName))
	if value == "" {
		return defaultValue
	}
	return value
}
