package main

import (
	"fmt"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func testCmd(ctx anpan.Context, _ []string) {
	ctx.Session.ChannelMessageSend(ctx.Channel.ID, "hello, world")
}
func pingCmd(ctx anpan.Context, _ []string) {
	ping, _ := ctx.Reply("Let me check that for you!")

	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne)

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf("Ping took `%s`!", took.String()),
	}

	ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, ping.ID, embed)
}
