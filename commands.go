package main

import (
	"github.com/bwmarrin/discordgo"
)

func pingCmd(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	s.ChannelMessageSend(m.ChannelID, "hello, world")
}
