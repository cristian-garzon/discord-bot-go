package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
)

type Router struct {
	*CommandsHandler
}

type HandlerFunc func(interface{})

func ConfigureRoutes(session *discordgo.Session, handler *CommandsHandler) {

	router := &Router{handler}
	session.AddHandler(
		func(session *discordgo.Session, message *discordgo.MessageCreate) {

			if message.Author.ID == session.State.User.ID {
				return
			}

			ctx := context.WithValue(context.Background(), "session", session)
			ctx = context.WithValue(ctx, "message", message)

			router.On("C! ping", ctx, handler.Ping)
			router.On("C! avatar", ctx, handler.Avatar)
			return
		},
	)
}

func (router *Router) On(command string, ctx context.Context, callback func(ctx context.Context)) {

	message := ctx.Value("message").(*discordgo.MessageCreate)

	if message.Content == command {
		callback(ctx)
	}
}
