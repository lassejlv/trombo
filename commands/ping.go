package commands

import "github.com/bwmarrin/discordgo"

func init() {
	Commands = append(Commands, Command{
		Name:        "ping",
		Description: "pong",
		Handler: func(s *discordgo.Session, m *discordgo.MessageCreate) {
			s.ChannelMessageSend(m.ChannelID, "pong")
		},
	})
}
