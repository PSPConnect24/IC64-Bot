package main

import (
	"os"
	"reflect"
	"strings"

	"github.com/MikeModder/anpan"
	"github.com/apaxa-go/eval"
)

func evalcmd(ctx anpan.Context, content []string) {
	exp, err := eval.ParseString(strings.Join(content, " "), "")
	if err != nil {
		ctx.Reply(":x: An error occurred: **" + err.Error() + "**")
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

	output, err := exp.EvalToRegular(env)
	if err != nil {
		ctx.Reply(":x: An error occurred: **" + err.Error() + "**")
		return
	}

	ctx.Reply(":white_check_mark: Output:\n```\n" + output.String() + "\n```")
}

func shutdowncmd(ctx anpan.Context, _ []string) {
	ctx.Session.MessageReactionAdd(ctx.Channel.ID, ctx.Message.ID, ":white_check_mark:")
	ctx.Session.Close()
	os.Exit(0)
}
