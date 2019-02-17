package main

import (
	"strings"

	gotag "github.com/Apfel/gotag"
	gotaglib "github.com/Apfel/gotag/libraries"
	"github.com/MikeModder/anpan"
)

var parser = gotag.NewBuilder().AddMethods(gotaglib.Arguments()).Build()

func tagcmd(ctx anpan.Context, args []string) {
	ctx.Reply(parser.Parse(strings.Join(args, " ")))
}
