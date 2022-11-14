package handlers

import (
	"discord-bot-go/dependencies"
	"github.com/bwmarrin/discordgo"
)

type GetAvatarHandler struct {
	GetAvatar dependencies.GetAvatar
}

func NewGetAvatarHandler(avatar dependencies.GetAvatar) *GetAvatarHandler {
	return &GetAvatarHandler{
		GetAvatar: avatar,
	}
}

func (Avatar *GetAvatarHandler) Execute(message *discordgo.MessageCreate) {
	Avatar.GetAvatar(message)
}
