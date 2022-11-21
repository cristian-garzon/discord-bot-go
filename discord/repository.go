package discord

import (
	"discord-bot-go/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Repository struct {
	sendMessage      func(m *discordgo.MessageCreate, message string)
	closeConnection  func()
	openConnection   func() error
	SendMessageEmbed func(m *discordgo.MessageCreate, embeds []*discordgo.MessageEmbed)
}

func NewRepository(client *BotClient) *Repository {
	return &Repository{
		sendMessage:      client.SendMessage,
		closeConnection:  client.CloseConnection,
		openConnection:   client.OpenConnection,
		SendMessageEmbed: client.SendMessageEmbed,
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
	Repository.SendMessageEmbed(m, embeds)
}

func (Repository *Repository) Help(m *discordgo.MessageCreate) {
	helpCommands := ""
	for _, v := range utils.Commands {
		helpCommands += fmt.Sprintf("**%s**: %s. %s \n", v.Title, v.Description, v.Prefix)
	}
	var embeds = []*discordgo.MessageEmbed{
		&discordgo.MessageEmbed{
			Title:       "Help",
			Description: helpCommands,
			Footer: &discordgo.MessageEmbedFooter{
				Text: fmt.Sprintf("cualquier duda que tengas dicelo a mi desarrollador ^^"),
			},
			Color: utils.Red,
		},
	}
	Repository.SendMessageEmbed(m, embeds)
}

func (Repository *Repository) Close() {
	Repository.closeConnection()
}

func (Repository *Repository) Open() error {
	return Repository.openConnection()
}
