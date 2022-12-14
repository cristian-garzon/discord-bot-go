package dependencies

import (
	"github.com/bwmarrin/discordgo"
)

type SendMessage func(m *discordgo.MessageCreate, message string)

type Close func()

type Open func() error

type GetAvatar func(m *discordgo.MessageCreate)

type GetHelp func(m *discordgo.MessageCreate)
