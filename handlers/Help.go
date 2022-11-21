package handlers

import (
	"discord-bot-go/dependencies"
	"github.com/bwmarrin/discordgo"
)

type GetHelpHandler struct {
	GetHelp dependencies.GetHelp
}

func NewGetHelpHandler(help dependencies.GetHelp) *GetHelpHandler {
	return &GetHelpHandler{
		GetHelp: help,
	}
}

func (Help *GetHelpHandler) Execute(message *discordgo.MessageCreate) {
	Help.GetHelp(message)
}
