package commands

import (
	"os"

	"github.com/MikeModder/anpan"
)

// Shutdown command
func Shutdown(ctx anpan.Context, _ []string) {
	ctx.Session.MessageReactionAdd(ctx.Channel.ID, ctx.Message.ID, ":white_check_mark:")
	ctx.Session.Close()
	os.Exit(0)
}
