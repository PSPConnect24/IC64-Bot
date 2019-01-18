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
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: ctx.Session.State.User.AvatarURL(""),
		},
	})
}
