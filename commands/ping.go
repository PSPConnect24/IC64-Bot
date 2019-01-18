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
	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne).Seconds() * float64(1000)

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf(":ping_pong: Pong! Ping: **%f**!", took),
	}

	ctx.Session.ChannelMessageEditEmbed(ctx.Channel.ID, ping.ID, embed)
}
