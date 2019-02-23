package main

import (
	"fmt"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func pingcmd(ctx anpan.Context, _ []string) error {
	ping, _ := ctx.Reply("Let me check that for you!")
	tsOne, _ := ping.Timestamp.Parse()
	took := time.Now().Sub(tsOne).Seconds() * float64(1000)

	embed := &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf(":ping_pong: Pong! Ping: **%f**!", took),
	}

	ctx.Session.ChannelMessageEditEmbed(ctx.Channel.ID, ping.ID, embed)
}

func aboutcmd(ctx anpan.Context, _ []string) error {
	ctx.ReplyEmbed(&discordgo.MessageEmbed{
		Title:       "About IronConnect64-Bot",
		Description: "IronConnect64-Bot is a bot made specifically for IronConnect64, made to give you related information and content for IronConnect64.",
		URL:         "https://ic64.xyz/",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Session.State.User.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{{
			Name:  "Related content",
			Value: "**[Website](https://ic64.xyz)**\n**[IC64 on GitHub](https://github.com/IronConnect64)**\n**[IC64-Bot's repository](https://github.com/IronConnect64/IC64-Bot)**",
		}},
	})
}
