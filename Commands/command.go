package Commands

import (
	"discord-bot-go/dependencies"
	"discord-bot-go/handlers"
	"discord-bot-go/utils"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Command struct {
	pingHandler *handlers.PingHandler
}

func NewCommands(repositories *dependencies.Repositories) *Command {
	return &Command{
		pingHandler: handlers.NewPing(repositories.SendMessage),
	}
}

func (Command *Command) SetupCommands(
	s *discordgo.Session,
	m *discordgo.MessageCreate,
) {

	if m.Author.ID == s.State.User.ID || !strings.Contains(m.Content, utils.PrefixPath) {
		return
	}

	command := strings.Split(m.Content, utils.PrefixPath)

	if command[0] != "" {
		return
	}
	//add commands HERE
	for _, v := range utils.Ping {
		if command[1] == v {
			Command.pingHandler.Execute(m)
		}
	}
}
