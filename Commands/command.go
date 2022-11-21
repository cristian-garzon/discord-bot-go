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
	getAvatar   *handlers.GetAvatarHandler
	GetHelp     *handlers.GetHelpHandler
}

func NewCommands(repositories *dependencies.Repositories) *Command {
	return &Command{
		pingHandler: handlers.NewPing(repositories.SendMessage),
		getAvatar:   handlers.NewGetAvatarHandler(repositories.GetAvatar),
		GetHelp:     handlers.NewGetHelpHandler(repositories.GetHelp),
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
	if utils.Contains(utils.Ping, command[1]) {
		Command.pingHandler.Execute(m)
		return
	}

	if utils.Contains(utils.Avatar, command[1]) {
		Command.getAvatar.Execute(m)
		return
	}

	if utils.Contains(utils.Help, command[1]) {
		Command.GetHelp.Execute(m)
		return
	}

}
