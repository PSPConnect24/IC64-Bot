package commands

import (
	"fmt"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

// Ping command.
func Ping(ctx anpan.Context, _ []string) {
	ping, _ := ctx.Reply("Let me check that for you!")
	d, _ := time.ParseDuration("1ms")
	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne).Round(d).String()

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf(":ping_pong: Pong! Ping: **%s**!", took),
	}

	ctx.Session.ChannelMessageEditEmbed(ctx.Message.ChannelID, ping.ID, embed)
}
