package commands

import (
	"reflect"
	"strings"

	"github.com/MikeModder/anpan"
	"github.com/apaxa-go/eval"
)

// Eval command.
func Eval(ctx anpan.Context, content []string) {
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
