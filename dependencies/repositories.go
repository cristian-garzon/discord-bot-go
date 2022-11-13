package dependencies

import (
	"discord-bot-go/discord"
)

type Repositories struct {
	SendMessage SendMessage
	Close       Close
	Open        Open
}

func NewRepositories(
	DiscordBot *discord.Discord,
) *Repositories {

	discordRepository := discord.NewRepository(DiscordBot.NewBotClient())

	return &Repositories{
		SendMessage: discordRepository.CreateMessage,
		Close:       discordRepository.Close,
		Open:        discordRepository.Open,
	}
}
