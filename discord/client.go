package discord

import (
	"fmt"
	discordgo "github.com/bwmarrin/discordgo"
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

func (client *BotClient) GetAvatar(m *discordgo.MessageCreate, embeds []*discordgo.MessageEmbed) {
	_, err := client.Bot.ChannelMessageSendEmbeds(m.ChannelID, embeds)
	if err != nil {
		client.Bot.ChannelMessageSend(m.ChannelID,
			fmt.Sprintf("Hubo un error mostrando tu avatar %s :c", m.Author.Mention()),
		)
	}
}

func (client *BotClient) CloseConnection() {
	client.Bot.Close()
}

func (client *BotClient) OpenConnection() error {
	return client.Bot.Open()
}
