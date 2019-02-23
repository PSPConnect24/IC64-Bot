package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func pingcmd(ctx anpan.Context, _ []string) error {
	ping, err := ctx.Reply("Let me check that for you!")
	if os.IsExist(err) {
		return err
	}

	tsOne, err := ping.Timestamp.Parse()
	if os.IsExist(err) {
		return err
	}

	took := time.Now().Sub(tsOne).Seconds() * float64(1000)
	_, err = ctx.Session.ChannelMessageEditEmbed(ctx.Channel.ID, ping.ID, &discordgo.MessageEmbed{
		Title:       "Pong!",
		Description: fmt.Sprintf(":ping_pong: Pong! Ping: **%f**!", took),
	})
	return err
}

func aboutcmd(ctx anpan.Context, _ []string) error {
	_, err := ctx.ReplyEmbed(&discordgo.MessageEmbed{
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

	return err
}
