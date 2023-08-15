package actions

import "github.com/bwmarrin/discordgo"

type Ping struct{}

func NewPing() (*Ping, error) {
	return &Ping{}, nil
}

func (action *Ping) Execute(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, "Pong!")
}
