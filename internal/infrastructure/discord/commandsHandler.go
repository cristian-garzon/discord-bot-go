package discord

import (
	"context"
	"doscord-bot-go/internal/infrastructure/injections"
	"github.com/bwmarrin/discordgo"
)

type CommandsHandler struct {
	*injections.ActionsBuilder
}

func NewCommandsHandler(builder *injections.ActionsBuilder) *CommandsHandler {
	return &CommandsHandler{builder}
}

func (handler *CommandsHandler) Ping(ctx context.Context) {

	message := ctx.Value("message").(*discordgo.MessageCreate)
	session := ctx.Value("session").(*discordgo.Session)
	handler.ActionsBuilder.Ping(session, message)

}

func (handler *CommandsHandler) Avatar(ctx context.Context) {

}
