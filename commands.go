package main

import (
	"github.com/bwmarrin/discordgo"
)

func testCmd(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) {
	s.ChannelMessageSend(m.ChannelID, "hello, world")
}
