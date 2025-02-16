package commands

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name        string
	Description string
	Handler     func(*discordgo.Session, *discordgo.MessageCreate)
}

var Commands []Command
