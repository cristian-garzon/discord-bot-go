package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type Repository struct {
	sendMessage     func(m *discordgo.MessageCreate, message string)
	closeConnection func()
	openConnection  func() error
}

func NewRepository(client *BotClient) *Repository {
	return &Repository{
		sendMessage:     client.SendMessage,
		closeConnection: client.CloseConnection,
		openConnection:  client.OpenConnection,
	}
}

func (Repository *Repository) CreateMessage(m *discordgo.MessageCreate, message string) {
	Repository.sendMessage(m, fmt.Sprintf("me llamas **<@%s>**?", m.Author.ID))
}

func (Repository *Repository) Close() {
	Repository.closeConnection()
}

func (Repository *Repository) Open() error {
	return Repository.openConnection()
}
