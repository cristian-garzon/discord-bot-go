package discord

import (
	"github.com/bwmarrin/discordgo"
)

type BotClient struct {
	Bot *discordgo.Session
}

func newBotClient(bot *discordgo.Session) *BotClient {
	return &BotClient{
		Bot: bot,
	}
}

func (client *BotClient) SendMessage(m *discordgo.MessageCreate, message string) {
	client.Bot.ChannelMessageSend(m.ChannelID, message)

}

func (client *BotClient) CloseConnection() {
	client.Bot.Close()
}

func (client *BotClient) OpenConnection() error {
	return client.Bot.Open()
}
