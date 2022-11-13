package dependencies

import (
	"github.com/bwmarrin/discordgo"
)

type SendMessage func(m *discordgo.MessageCreate, message string)

type Close func()

type Open func() error
