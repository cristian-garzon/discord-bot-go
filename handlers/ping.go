package handlers

import (
	"discord-bot-go/dependencies"
	"github.com/bwmarrin/discordgo"
)

type PingHandler struct {
	SendMessage dependencies.SendMessage
}

func NewPing(sendMessage dependencies.SendMessage) *PingHandler {
	return &PingHandler{
		SendMessage: sendMessage,
	}
}

func (ping *PingHandler) Execute(message *discordgo.MessageCreate) {
	ping.SendMessage(message, "me llamas?")
}
