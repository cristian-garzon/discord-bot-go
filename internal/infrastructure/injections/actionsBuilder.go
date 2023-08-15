package injections

import (
	"doscord-bot-go/internal/domain/actions"
	"github.com/bwmarrin/discordgo"
)

type ActionsBuilder struct {
	Ping func(session *discordgo.Session, message *discordgo.MessageCreate)
}

func NewActionsBuilder() (*ActionsBuilder, error) {
	ping, err := actions.NewPing()

	if err != nil {
		return nil, err
	}

	return &ActionsBuilder{
		ping.Execute,
	}, err
}
