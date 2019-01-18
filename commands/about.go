package commands

import (
	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

// About command.
func About(ctx anpan.Context, _ []string) {
	ctx.ReplyEmbed(&discordgo.MessageEmbed{
		Title:       "About IronConnect64-Bot",
		Description: "IronConnect64-Bot is a bot made specifically for IronConnect64, made to give you related information and content for IronConnect64.",
		URL:         "https://ic64.xyz/",
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Session.State.User.AvatarURL(""),
		},
		Fields: []*discordgo.MessageEmbedField{&discordgo.MessageEmbedField{
			Name:  "Related content",
			Value: "**[Website](https://ic64.xyz)**\n**[IC64 on GitHub](https://github.com/IronConnect64)**\n**[IC64-Bot's repository](https://github.com/IronConnect64/IC64-Bot)**",
		}},
	})
}
