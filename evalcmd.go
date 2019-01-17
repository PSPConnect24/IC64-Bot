package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/MikeModder/anpan"
	"github.com/apaxa-go/eval"
)

func evalCmd(ctx anpan.Context, content []string) {
	exp, err := eval.ParseString(strings.Join(content, " "), "")
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Channel.ID, ":x: An error occurred: **"+err.Error()+"**")
		return
	}

	env := eval.Args{
		"session": eval.MakeDataRegular(reflect.ValueOf(ctx.Session)),
		"channel": eval.MakeDataRegular(reflect.ValueOf(ctx.Channel)),
		"user":    eval.MakeDataRegular(reflect.ValueOf(ctx.User)),
		"message": eval.MakeDataRegular(reflect.ValueOf(ctx.Message)),
		"guild":   eval.MakeDataRegular(reflect.ValueOf(ctx.Guild)),
		"bot":     eval.MakeDataRegular(reflect.ValueOf(ctx.Session.State)),
	}

	output, err := exp.EvalToInterface(env)
	if err != nil {
		ctx.Session.ChannelMessageSend(ctx.Channel.ID, ":x: An error occurred: "+err.Error())
		return
	}

	ctx.Session.ChannelMessageSend(ctx.Channel.ID, ":white_check_mark: Output:\n```\n"+fmt.Sprintf("%d", output)+"\n```")
}

func shutdownCmd(ctx anpan.Context, _ []string) {
	ctx.Session.MessageReactionAdd(ctx.Channel.ID, ctx.Message.ID, ":white_check_mark:")
	ctx.Session.Close()
	os.Exit(0)
}
