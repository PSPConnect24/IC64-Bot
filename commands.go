package main

import (
	"github.com/MikeModder/anpan"
)

func testCmd(ctx anpan.Context, _ []string) {
	ctx.Session.ChannelMessageSend(ctx.Channel.ID, "hello, world")
}
