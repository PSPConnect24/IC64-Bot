package main

import (
	"fmt"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

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

func aboutCmd(ctx anpan.Context, _ []string) {
	ctx.ReplyEmbed(&discordgo.MessageEmbed{
		Title:       "About IronConnect64-Bot",
		Description: "IronConnect64-Bot is a bot made specifically for IronConnect64, made to give you related information and content for IronConnect64.",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Session.State.User.AvatarURL(""),
		},
	})
}
