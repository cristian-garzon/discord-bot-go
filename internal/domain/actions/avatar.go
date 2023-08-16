package actions

import (
	"github.com/bwmarrin/discordgo"
)

type Avatar struct{}

func NewAvatar() (*Avatar, error) {
	return &Avatar{}, nil
}

func (action *Avatar) Execute(session *discordgo.Session, message *discordgo.MessageCreate) {
	avatarUrl := message.Author.AvatarURL("2048")
	if avatarUrl != "" {
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title: "avatar",
			Image: &discordgo.MessageEmbedImage{
				URL: avatarUrl,
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "esta bonita tu foto " + message.Author.Username,
			},
			Color: 15548997,
		})
	} else {
		session.ChannelMessageSend(message.ChannelID, "error al obtener el avatar :c")
	}
}
