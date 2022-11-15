package discord

import (
	"discord-bot-go/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Repository struct {
	sendMessage     func(m *discordgo.MessageCreate, message string)
	closeConnection func()
	openConnection  func() error
	getAvatar       func(m *discordgo.MessageCreate, embeds []*discordgo.MessageEmbed)
}

func NewRepository(client *BotClient) *Repository {
	return &Repository{
		sendMessage:     client.SendMessage,
		closeConnection: client.CloseConnection,
		openConnection:  client.OpenConnection,
		getAvatar:       client.GetAvatar,
	}
}

func (Repository *Repository) CreateMessage(m *discordgo.MessageCreate, message string) {
	Repository.sendMessage(m, fmt.Sprintf("me llamas **<@%s>**?", m.Author.ID))
}

func (Repository *Repository) Avatar(m *discordgo.MessageCreate) {
	var embeds = []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Title:       "Avatar",
			Description: fmt.Sprintf("Avatar de %s", m.Author.Mention()),
			Image: &discordgo.MessageEmbedImage{
				URL:    m.Author.AvatarURL("1024"),
				Height: 200,
				Width:  200,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("me agrada tu avatar :D"),
			},
			Color: utils.Yellow,
		},
	}
	Repository.getAvatar(m, embeds)
}

func (Repository *Repository) Close() {
	Repository.closeConnection()
}

func (Repository *Repository) Open() error {
	return Repository.openConnection()
}
