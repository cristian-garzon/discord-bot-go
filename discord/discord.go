package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Discord struct {
	DiscordBot *discordgo.Session
}

func NewDiscord(config Config) (*Discord, error) {
	botToken := fmt.Sprintf("Bot %s", config.Token)

	bot, err := discordgo.New(botToken)

	if err != nil {
		return nil, err
	}

	bot.Identify.Intents = discordgo.IntentGuildMessages

	return &Discord{
		DiscordBot: bot,
	}, nil
}

func (bot *Discord) NewBotClient() *BotClient {
	return newBotClient(bot.DiscordBot)
}
